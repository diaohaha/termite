package scheduler

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/stable"
	"fmt"
	"github.com/go-errors/errors"
	"go.uber.org/zap"
	"sync"
	"time"
)

func FlowRecover(partitions []int) {
	// 修复上失败的flow
	flows, _ := db.GetErrorFlowsByPartitions(partitions)
	for _, flow := range flows {
		err := db.RecoverWorkFlow(flow.Id)
		if err != nil {
			stable.SchedulerLogger.Error("FlowRecover Error!", zap.Error(err))
		}
	}
	return
}

func FlowModify(partitions []int) {
	// TODO: 减少数据库查询
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if perr := recover(); perr != nil {
			stable.SchedulerLogger.Error("FlowModify panic", zap.Error(perr.(error))) // 这里的err其实就是panic传入的内容
			stable.CaptureError(perr.(error), "FlowModify", "FlowModify", map[string]string{}, map[string]string{
				"method": "FlowModify",
				"type":   "panic",
			})
			time.Sleep(1 * time.Second)
			return
		}
	}()
	length := 1000
	var sid = int64(0)
	for flowName, flowConfig := range mem.GlobalTermiteConfig.FlowConfigMap {
		// 1.delay running 全量检查
		stime := time.Now().UnixNano()
		fmt.Println("工作流状态同步:", flowName)
		subStime := time.Now().UnixNano()
		for {
			flows := db.GetFlowsByPartitionsAndVflow(sid, partitions, flowName, []int64{
				dal.TermiteWorkFlowStateRunning,
				dal.TermiteWorkFlowStateDelay,
			}, length)
			if len(flows) == 0 {
				break
			}
			for _, flow := range flows {
				simpleFlowModifyCore(flow)
				sid = flow.Id
			}
		}
		subEtime := time.Now().UnixNano()
		fmt.Println("工作流状态同步:", flowName, "执行中&延期任务同步: 耗时:", (subEtime-subStime)/int64(time.Millisecond), "ms")

		// 2. init 限量检查
		if flowConfig.Config.MaxSchedulerCount == 0 {
			// 全量检查
			sid = int64(0)
			subStime := time.Now().UnixNano()
			for {
				flows := db.GetFlowsByPartitionsAndVflow(sid, partitions, flowName, []int64{
					dal.TermiteWorkFlowStateInit,
				}, length)
				if len(flows) == 0 {
					break
				}
				for _, flow := range flows {
					simpleFlowModifyCore(flow)
					sid = flow.Id
				}
			}
			subEtime := time.Now().UnixNano()
			fmt.Println("工作流状态同步:", flowName, "待调度任务全量同步: 耗时:", (subEtime-subStime)/int64(time.Millisecond), "ms")
		} else {
			// 限量检查
			subStime := time.Now().UnixNano()
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
			}, flowConfig.Config.MaxSchedulerCount, order)
			for _, flow := range flows {
				simpleFlowModifyCore(flow)
				sid = flow.Id
			}
			subEtime := time.Now().UnixNano()
			fmt.Println("工作流状态同步:", flowName, "待调度任务限量同步: 限量", flowConfig.Config.MaxSchedulerCount, "耗时:", (subEtime-subStime)/int64(time.Millisecond), "ms")
		}
		etime := time.Now().UnixNano()
		fmt.Println("工作流状态同步:", flowName, "耗时:", (etime-stime)/int64(time.Millisecond), "ms")
	}

	//for {
	//	//var wg = &sync.WaitGroup{}
	//	//ch := make(chan bool, 5)
	//	flows := db.GetFlowsByPartitions(sid, partitions, vstates, length)
	//	if len(flows) == 0 {
	//		break
	//	}
	//	stable.SchedulerLogger.Info(fmt.Sprintf("modify flows count: %d", len(flows)))
	//	for _, flow := range flows {
	//		//wg.Add(1)
	//		//go flowModifyCore(ch, wg, flow)
	//		simpleFlowModifyCore(flow)
	//		sid = flow.Id
	//	}
	//	//time.Sleep(time.Millisecond * 5)
	//	//wg.Wait()
	//}
	return
}

func simpleFlowModifyCore(flow dal.TermiteFlow) {
	works := db.GetWorksByFlowAndCid(flow.Vflow, flow.Cid)
	if len(works) == 0 {
		_ = db.DeleteFlow(flow.Id)
		return
	}
	flowState, err := getFlowStateByWorksState(works)
	if err == nil {
		err = db.UpdateFlowState(flow.Id, flowState)
		if err != nil {
			stable.SchedulerLogger.Error("simpleFlowModifyCore Error!", zap.Error(err))
		}
	}
	if flowState == dal.TermiteWorkFlowStateFinish {
		err = db.DeleteFlow(flow.Id)
		if err != nil {
			stable.SchedulerLogger.Error("simpleFlowModifyCore Error!", zap.Error(err))
		}
	}
}

func flowModifyCore(ch chan bool, wg *sync.WaitGroup, flow dal.TermiteFlow) {
	defer wg.Done()
	ch <- true
	stime := time.Now().UnixNano()
	works := db.GetWorksByFlowAndCid(flow.Vflow, flow.Cid)
	if len(works) == 0 {
		_ = db.DeleteFlow(flow.Id)
		<-ch
		return
	}
	flowState, err := getFlowStateByWorksState(works)
	if err == nil {
		err = db.UpdateFlowState(flow.Id, flowState)
		if err != nil {
			stable.SchedulerLogger.Error("flowModifyCore Error!", zap.Error(err))
		}
	}
	if flowState == dal.TermiteWorkFlowStateFinish {
		err = db.DeleteFlow(flow.Id)
		if err != nil {
			stable.SchedulerLogger.Error("flowModifyCore Error!", zap.Error(err))
		}
	}
	etime := time.Now().UnixNano()
	stable.SchedulerLogger.Info("耗时:", zap.Int64("耗时", (etime-stime)/int64(time.Millisecond)))
	<-ch
}

func getFlowStateByWorksState(works []dal.TermiteWork) (int, error) {
	// 考虑工作流配置减少work配置时的兼容
	c := 0
	init_num := 0
	finish_num := 0
	is_running := false
	is_timeout := false
	is_dealy := false
	for _, work := range works {
		c += 1
		switch work.Vstate {
		case dal.TermiteWorkStateInit:
			init_num += 1
		case dal.TermiteWorkStatePush:
			is_running = true
		case dal.TermiteWorkStateRunning:
			is_running = true
		case dal.TermiteWorkStateDelay:
			is_dealy = true
		case dal.TermiteWorkStateError:
			return dal.TermiteWorkFlowStateError, nil
		case dal.TermiteWorkStateTimeout:
			is_timeout = true
		case dal.TermiteWorkStateFailed:
			finish_num += 1
			return dal.TermiteWorkFlowStateFailed, nil
		case dal.TermiteWorkStateSuccess:
			finish_num += 1
		}
	}
	if is_dealy {
		return dal.TermiteWorkFlowStateDelay, nil
	}
	if is_timeout {
		return dal.TermiteWorkFlowStateTimeout, nil
	}
	if is_running {
		return dal.TermiteWorkFlowStateRunning, nil
	}
	if c == finish_num {
		return dal.TermiteWorkFlowStateFinish, nil
	}
	// 考虑工作流配置减少work配置时的兼容
	if len(mem.GlobalTermiteConfig.FlowConfigMap[works[0].Vflow].Config.Works) == finish_num {
		return dal.TermiteWorkFlowStateFinish, nil
	}
	if init_num == c {
		return dal.TermiteWorkFlowStateInit, nil
	}
	if (0 < init_num) && (init_num < c) {
		return dal.TermiteWorkFlowStateRunning, nil
	}
	return -1, errors.New("unknow flow state.")
}
