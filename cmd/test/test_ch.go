package main

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/scheduler"
	"log"
	"sync"
	"time"
)

func main() {
	//wg := &sync.WaitGroup{}
	//ch := make(chan bool, 10)
	//values := []int{1, 2, 3, 4}
	//for _, value := range values {
	//	wg.Add(1)
	//	go printValue(ch, wg, value)
	//}
	//time.Sleep(time.Microsecond * 10)
	//wg.Wait()
	db.InitDB()
	mem.InitMem()
	scheduler.Init()
	//newDag, _ := dag.DeepCopyDagNode(*scheduler.GlobalflowDagGraphMap["video_h265"], map[string]*model.DagNode{})
	//dagNew := *scheduler.GlobalflowDagGraphMap["video_h265"]
	//dag.PrintNode(&newDag)
	//dag.PrintNode(&dagNew)
	//dag.PrintNode(scheduler.GlobalflowDagGraphMap["video_h265"])
	//for _, dagNode := range scheduler.GlobalflowDagGraphMap {
	//	dag.PrintNode(*dagNode)
	//}
	var paritions []int
	for i := 0; i <= 127; i++ {
		paritions = append(paritions, i)
	}
	scheduler.FlowRecover(paritions)
}

func printValue(ch chan bool, wg *sync.WaitGroup, v int) {
	defer wg.Done()
	ch <- true
	time.Sleep(time.Second * time.Duration(v))
	log.Println(v)
	<-ch
}
