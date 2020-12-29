package scheduler

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/stable"
	"fmt"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
	Support multiple dag scheduling nodes
*/

type SchedulerNode struct {
	NodeType   string
	NodeId     string
	Partitions []int
}

var Inode SchedulerNode
var liveNodeCount int
var expireTime time.Time

func setNodeId() {
	// 节点ID
	//cmd := exec.Command("cat", "/etc/machine-id")
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	stable.Logger.Error("Get Node Id Error", zap.Error(err), zap.String("err", stderr.String()))
	//	panic(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String())
	//NodeId = out.String()

	// 使用uuid当节点runtime id
	if Inode.NodeId == "" {
		Inode.NodeId = uuid.NewV4().String()
	}
}

func SetNodeType(nodeType string) {
	Inode.NodeType = nodeType
}

func HeartBeat() {
	setNodeId()
	for {
		log.Println("Node Hearbeat ...")
		err := db.NodeHeartBeat(Inode.NodeId, Inode.NodeType)
		if err != nil {
			stable.SchedulerLogger.Error("heartbeat error", zap.Error(err))
			panic(err)
		}
		time.Sleep(5 * time.Second)
	}
}

func DeleteExpireNodes() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			stable.SchedulerLogger.Error("DeleteExpireNodes panic", zap.Error(err.(error))) // 这里的err其实就是panic传入的内容
			stable.CaptureError(err.(error), "DeleteExpireNodes", "DeleteExpireNodes", map[string]string{}, map[string]string{
				"method": "DeleteExpireNodes",
				"type":   "panic",
			})
			return
		}
	}()
	fmt.Println("DeleteExpireNodes: start...")
	err := db.DeleteExpireNodes()
	if err != nil {
		stable.CaptureError(err.(error), "DeleteExpireNodes", "DeleteExpireNodes", map[string]string{}, map[string]string{
			"method": "db.DeleteExpireNodes",
		})
	}
	fmt.Println("DeleteExpireNodes: end...")
}

func Rebalance() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			stable.SchedulerLogger.Error("Rebalance panic", zap.Error(err.(error))) // 这里的err其实就是panic传入的内容
			stable.CaptureError(err.(error), "Rebalance", "Rebalance", map[string]string{}, map[string]string{
				"method": "Rebalance",
				"type":   "panic",
			})
			time.Sleep(1 * time.Second) // sentry 上报
			return
		}
	}()
	// 只有dag进行多节点调度 删除过期节点
	log.Println("reblance start...")
	nodes, err := db.GetLiveNodes(dal.NODETYPE_DAG)
	if err != nil {
		stable.CaptureError(err.(error), "Rebalance", "Rebalance", map[string]string{}, map[string]string{
			"method": "GetLiveNodes",
		})
		return
	}
	nodeMap := map[string]*SchedulerNode{}
	for _, node := range nodes {
		nodeMap[node.NodeId] = &SchedulerNode{
			Partitions: []int{},
			NodeId:     node.NodeId,
			NodeType:   node.NodeType,
		}
	}

	// 重新分配分区
	partition := 0
	var runSign bool
	for {
		runSign = true
		for _, node := range nodeMap {
			node.Partitions = append(node.Partitions, partition)
			partition = partition + 1
			if partition >= 128 {
				runSign = false
				break
			}
		}
		if !runSign {
			break
		}
	}
	log.Println("rebalance result: ", nodeMap)

	for _, node := range nodeMap {
		partitionStrList := []string{}
		for _, partition := range node.Partitions {
			partitionStrList = append(partitionStrList, strconv.Itoa(partition))

		}
		partitionsStr := strings.Join(partitionStrList, ",")
		err = db.UpdateNodePartitions(node.NodeId, partitionsStr)
		if err != nil {
			stable.CaptureError(err.(error), "Rebalance", "Rebalance", map[string]string{}, map[string]string{
				"method": "UpdateNodePartitions",
			})
		}
	}
	log.Println("reblance end...")
}

func GetLiveNodeCount() (count int, err error) {
	// 返回在线的count数量
	if time.Now().After(expireTime) {
		// refresh
		nodes, err := db.GetLiveNodes(dal.NODETYPE_DAG)
		if err != nil {
			stable.CaptureError(err.(error), "GetLiveNodeCount", "GetLiveNodeCount", map[string]string{}, map[string]string{
				"method": "GetLiveNodes",
			})
		}
		expireTime = time.Now().Add(100 * time.Second)
		liveNodeCount = len(nodes)
	}
	return liveNodeCount, err
}

func GetMyParitions() (partitions []int, err error) {
	node, err := db.QueryTermiteNodes(Inode.NodeId)
	if err != nil {
		stable.CaptureError(err.(error), "GetMyParitions", "GetMyParitions", map[string]string{}, map[string]string{
			"method": "QueryTermiteNodes",
		})
		return
	}
	if node.Partitions == "" {
		return
	}
	partitionStrList := strings.Split(node.Partitions, ",")
	for _, partitionStr := range partitionStrList {
		partitionInt, err := strconv.Atoi(partitionStr)
		if err != nil {
			stable.CaptureError(err.(error), "GetMyParitions", "GetMyParitions", map[string]string{}, map[string]string{
				"method": "Atoi",
			})
			panic(err)
		}
		partitions = append(partitions, partitionInt)
	}
	return
}

//func MasterDeamon() {
//	// 主节点任务 cron 充当
//	for {
//		time.Sleep(5 * time.Second)
//		Rebalance()
//	}
//}
