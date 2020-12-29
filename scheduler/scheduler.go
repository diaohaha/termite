package scheduler

import (
	"github.com/diaohaha/termite/dag/model"
	Dag "github.com/diaohaha/termite/dag/model"
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/dal/mq"
	"github.com/diaohaha/termite/stable"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type FlowDagGraphMap map[string]*model.DagNode

var flowDagGraphMap FlowDagGraphMap
var GlobalflowDagGraphMap FlowDagGraphMap
var stateMap *map[int]int

func Init() {
	// 任务 & DAG 状态MAP
	stateMap = &map[int]int{
		dal.TermiteWorkStateInit:    Dag.STATE_WAITING,
		dal.TermiteWorkStatePush:    Dag.STATE_RUNING,
		dal.TermiteWorkStateRunning: Dag.STATE_RUNING,
		dal.TermiteWorkStateSuccess: Dag.STATE_FINISH,
		dal.TermiteWorkStateFailed:  Dag.STATE_RUNING,
		dal.TermiteWorkStateError:   Dag.STATE_RUNING,
		dal.TermiteWorkStateTimeout: Dag.STATE_RUNING,
		dal.TermiteWorkStateDelay:   Dag.STATE_RUNING,
	}
	GlobalflowDagGraphMap = GerenateDagMap()
	// 初始化DAG图
	go func() {
		for true {
			time.Sleep(time.Second * 30)
			GlobalflowDagGraphMap = GerenateDagMap()
		}
	}()
}

func GerenateDagMap() FlowDagGraphMap {
	flowDagGraphMap = make(FlowDagGraphMap)
	mem.TryRefresh()
	for flow, flowConfig := range mem.GlobalTermiteConfig.FlowConfigMap {
		var workList = map[string]int{}
		tmp := []Dag.DagConfig{}
		for work, dagConfig := range flowConfig.Config.Dags {
			tmp = append(tmp, Dag.DagConfig{Name: work, Dependences: dagConfig.Dependences})
			workList[work] = 1
		}
		for _, work := range flowConfig.Config.Works {
			if _, ok := workList[work]; ok {
				continue
			} else {
				tmp = append(tmp, Dag.DagConfig{Name: work, Dependences: []string{}})
			}
		}
		flowDagGraphMap[flow] = Dag.InitDag(&tmp)
	}
	return flowDagGraphMap
}

func DagScheduler(partitions []int) {
	/*  DAG调度  */

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if perr := recover(); perr != nil {
			stable.SchedulerLogger.Error("DagScheduler panic", zap.Error(perr.(error))) // 这里的err其实就是panic传入的内容
			stable.CaptureError(perr.(error), "DagScheduler", "DagScheduler", map[string]string{}, map[string]string{
				"method": "DagScheduler",
				"type":   "panic",
			})
			time.Sleep(1 * time.Second)
			return
		}
	}()

	sid := int64(0)
	lenth := 1000
	var vstates []int64
	for flowName, flowConfig := range mem.GlobalTermiteConfig.FlowConfigMap {
		if flowConfig.Switch == 0 {
			fmt.Println("调度工作流:", flowName, "开关关闭")
			continue
		}
		stime := time.Now().UnixNano()
		raiseUpCount := 0
		runningCount := 0

		subStime := time.Now().UnixNano()
		for {
			// running状态全调度
			vstates = []int64{
				dal.TermiteWorkFlowStateRunning,
			}
			flows := db.GetFlowsByPartitionsAndVflow(sid, partitions, flowName, vstates, lenth)
			if len(flows) == 0 {
				break
			}
			for _, flow := range flows {
				runningCount = runningCount + 1
				sid = flow.Id
				c := simpleDagCore(flow, GlobalflowDagGraphMap)
				if c != 0 {
					raiseUpCount = raiseUpCount + 1
				}
			}
		}
		subEtime := time.Now().UnixNano()
		fmt.Println("调度工作流:", flowName, "running调度耗时: ", (subEtime-subStime)/int64(time.Millisecond), "ms")
		time.Sleep(time.Millisecond * 5)
		fmt.Println("调度工作流:", flowName, "执行中工作流数量: ", runningCount, "调度拉起工作流数量: ", raiseUpCount)
		if flowConfig.Config.MaxSchedulerCount == 0 {
			// 默认全量调度
			subStime := time.Now().UnixNano()
			fmt.Println("调度工作流:", flowName, "全量调度")
			sid = 0
			raiseUpCount := 0
			schedulerCount := 0
			for {
				flows := db.GetFlowsByPartitionsAndVflow(sid, partitions, flowName, []int64{
					dal.TermiteWorkFlowStateInit,
				}, lenth)
				if len(flows) == 0 {
					break
				}
				for _, flow := range flows {
					schedulerCount = schedulerCount + 1
					sid = flow.Id
					c := simpleDagCore(flow, GlobalflowDagGraphMap)
					if c != 0 {
						raiseUpCount = raiseUpCount + 1
					}
				}
			}
			subEtime := time.Now().UnixNano()
			fmt.Println("调度工作流:", flowName, "全量调度耗时: ", (subEtime-subStime)/int64(time.Millisecond), "ms")
			fmt.Println("调度工作流:", flowName, "待调度工作流数量: ", schedulerCount, "调度拉起工作流数量: ", raiseUpCount)
		} else {
			subStime := time.Now().UnixNano()
			fmt.Println("调度工作流:", flowName, "限量调度")
			count := 0 // 调度余量
			nodeCount, err := GetLiveNodeCount()
			if err != nil {
				stable.CaptureError(err, "DagScheduler", "DagScheduler", map[string]string{}, map[string]string{
					"method":    "DagScheduler",
					"msg":       "GetLiveNodeCount() Error!",
					"flow_name": flowName,
				})
			}
			SchedulingCount, err := db.GetFlowCountByStates(flowName, []int64{
				dal.TermiteWorkFlowStateRunning,
				dal.TermiteWorkFlowStateDelay,
			})
			if err != nil {
				stable.CaptureError(err, "DagScheduler", "DagScheduler", map[string]string{}, map[string]string{
					"method":    "DagScheduler",
					"msg":       "GetFlowCountByStates() Error!",
					"flow_name": flowName,
				})
			}
			fmt.Println("nodeCount:", nodeCount, "SchedulingCount:", SchedulingCount)
			count = (flowConfig.Config.MaxSchedulerCount - SchedulingCount) / nodeCount
			fmt.Println("调度工作流:", flowName, "调度余量:", count, " MaxSchedulerCount:", flowConfig.Config.MaxSchedulerCount)
			if count > 0 {
				var order string
				switch flowConfig.Config.SchedulerMode {
				case dal.SCHEDULER_MODE_FIFO:
					order = "id asc"
				case dal.SCHEDULER_MODE_LIFO:
					order = "id desc"
				default:
					order = "id asc"
				}
				flows := db.GetFlowsByPartitionsAndVflowOrdered(partitions, flowName, []int64{
					dal.TermiteWorkFlowStateInit,
				}, count, order)
				raiseUpCount := 0
				for _, flow := range flows {
					sid = flow.Id
					c := simpleDagCore(flow, GlobalflowDagGraphMap)
					if c != 0 {
						raiseUpCount = raiseUpCount + 1
					}
				}
				fmt.Println("调度工作流:", flowName, "待调度工作流数量: ", len(flows), "调度拉起工作流数量: ", raiseUpCount)
			}
			subEtime := time.Now().UnixNano()
			fmt.Println("调度工作流:", flowName, "限量调度耗时: ", (subEtime-subStime)/int64(time.Millisecond), "ms")
		}
		etime := time.Now().UnixNano()
		fmt.Println("调度工作流:", flowName, "耗时:", (etime-stime)/int64(time.Millisecond), "ms")
		//wg.Wait()

	}
	return
}

//func deepCopyFlowDagGraphMap(flowDagGraphMapSrc FlowDagGraphMap) (flowDagGraphMapDst FlowDagGraphMap) {
//	flowDagGraphMapDst = make(map[string]*model.DagNode)
//	for flow, flowDag := range flowDagGraphMapSrc {
//		//var dagNode model.DagNode
//		dagNode, _ := dag.DeepCopyDagNode(*flowDag, map[string]*Dag.DagNode{})
//		flowDagGraphMapDst[flow] = &dagNode
//	}
//	return
//}

func simpleDagCore(flow dal.TermiteFlow, flowDagGraphMapIns FlowDagGraphMap) (readyCount int) {
	//fmt.Printf("flowDagGraphMapIns Addr: %T Value:%v\n", flowDagGraphMapIns, flowDagGraphMapIns)
	var workStatus map[string]int
	var workIds map[string]int64
	//step2: 调度liveflow下的work
	works := db.GetWorksByFlowAndCid(flow.Vflow, flow.Cid)
	workStatus = map[string]int{}
	workIds = map[string]int64{}
	for _, work := range works {
		workStatus[work.Vwork] = (*stateMap)[int(work.Vstate)]
		workIds[work.Vwork] = work.Id
	}
	if len(workIds) == 0 {
		stable.SchedulerLogger.Info("flow exist, but work not exist.")
	} else {
		if _, ok := flowDagGraphMapIns[flow.Vflow]; ok {
			flowDagGraphMapIns[flow.Vflow].UpdateStatus(workStatus)
			isReadies := flowDagGraphMapIns[flow.Vflow].GetReadyNodes()
			readyCount = len(isReadies)
			stable.SchedulerLogger.Info("isReadise", zap.Int("ready num", len(isReadies)))
			for _, isReady := range isReadies {
				// 任务下发
				err := db.SetWorkPush(workIds[isReady])
				if err != nil {
					stable.CaptureError(err, "DagCore", "DagCore", map[string]string{}, map[string]string{})
					stable.SchedulerLogger.Error("err", zap.Error(err))
				} else {
					// TODO:事务
					mq.SendTaskReadyMsg(flow.Vflow, flow.Cid, isReady, flow.Project, workIds[isReady], flow.Id, mem.GlobalTermiteConfig.WorkConfigMap[isReady].Config)
					stable.SchedulerLogger.Info(fmt.Sprintf("flow:%s task:%s cid:%s workid:%d is ready",
						flow.Vflow, isReady, flow.Cid, workIds[isReady]))
				}
			}
		} else {
			// scheduler的flow未加载
			fmt.Println("Vflow:", flow.Vflow, "配置未加载！")
		}
	}
	return
}

//
//func DagCore(ch chan bool, wg *sync.WaitGroup, flow dal.TermiteFlow, flowDagGraphMapIns FlowDagGraphMap) {
//	defer wg.Done()
//	ch <- true
//	//fmt.Printf("flowDagGraphMapIns Addr: %T Value:%v\n", flowDagGraphMapIns, flowDagGraphMapIns)
//	var workStatus map[string]int
//	var workIds map[string]int64
//	//step2: 调度liveflow下的work
//	works := db.GetWorksByFlowAndCid(flow.Vflow, flow.Cid)
//	workStatus = map[string]int{}
//	workIds = map[string]int64{}
//	for _, work := range works {
//		workStatus[work.Vwork] = (*stateMap)[int(work.Vstate)]
//		workIds[work.Vwork] = work.Id
//	}
//	if len(workIds) == 0 {
//		stable.SchedulerLogger.Info("flow exist, but work not exist.")
//	} else {
//		flowDagGraphMapIns[flow.Vflow].UpdateStatus(workStatus)
//		isReadies := flowDagGraphMapIns[flow.Vflow].GetReadyNodes()
//		stable.SchedulerLogger.Info("isReadise", zap.Int("ready num", len(isReadies)))
//		for _, isReady := range isReadies {
//			// 任务下发
//			err := db.SetWorkPush(workIds[isReady])
//			if err != nil {
//				stable.CaptureError(err, "DagCore", "DagCore", map[string]string{}, map[string]string{})
//				stable.SchedulerLogger.Error("err", zap.Error(err))
//			} else {
//				mq.SendTaskReadyMsg(flow.Vflow, flow.Cid, isReady, workIds[isReady], flow.Id, mem.GlobalTermiteConfig.WorkConfigMap[isReady].Config)
//				stable.SchedulerLogger.Info(fmt.Sprintf("flow:%s task:%s cid:%s workid:%d is ready",
//					flow.Vflow, isReady, flow.Cid, workIds[isReady]))
//			}
//		}
//	}
//	<-ch
//}

func DelayScheduler() {
	// Delay任务调度 task粒度
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if perr := recover(); perr != nil {
			stable.SchedulerLogger.Error("DelayScheduler panic", zap.Error(perr.(error))) // 这里的err其实就是panic传入的内容
			stable.CaptureError(perr.(error), "DelayScheduler", "DelayScheduler", map[string]string{}, map[string]string{
				"method": "DelayScheduler",
				"type":   "panic",
			})
			time.Sleep(1 * time.Second)
			return
		}
	}()
	length := 2000
	sid := int64(0)
	for {
		stime := time.Now().UnixNano()
		tasks := db.GetTopDelayTasks(sid, length)
		if len(tasks) == 0 {
			break
		}

		for _, task := range tasks {
			// THINK: 这里多一次调用是不是有必要
			sid = task.Id
			flow, err := db.GetFlowByFlowKeyAndCid(task.Vflow, task.Cid)
			if err != nil {
				stable.SchedulerLogger.Error("GetFlowByFlowKeyAndCid Error", zap.Error(err))
				continue
			}
			if flow.Vstate == dal.TermiteWorkFlowStateError {
				// 异常任务停止delay任务
				continue
			}
			err = db.SetWorkPush(task.Id)
			if err != nil {
				stable.CaptureError(err, "DelayScheduler", "DelayScheduler", map[string]string{}, map[string]string{})
				stable.SchedulerLogger.Error("err", zap.Error(err))
			} else {
				mq.SendTaskReadyMsg(task.Vflow, task.Cid, task.Vwork, task.Project, task.Id, flow.Id, mem.GlobalTermiteConfig.WorkConfigMap[task.Vwork].Config)
				stable.SchedulerLogger.Info(fmt.Sprintf("flow:%s task:%s cid:%s workid:%d is ready",
					task.Vflow, task.Vwork, task.Cid, task.Id))
			}
		}
		etime := time.Now().UnixNano()
		fmt.Println("拉起", len(tasks), "个延期定时任务, 耗时:", (etime-stime)/int64(time.Millisecond), "ms")
	}
	return
}
