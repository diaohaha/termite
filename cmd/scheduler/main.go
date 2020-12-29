package main

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/scheduler"
	"github.com/diaohaha/termite/stable"
	"fmt"
	"go.uber.org/zap"
	"os"
	"runtime"

	"time"
)

//var pConfig rocketmq.ProducerConfig

func initLocal() {
	// sentry config
	//err := sentry.Init(sentry.ClientOptions{
	//	Dsn: dal.Env.Sentry_Dsn,
	//})
	//if err != nil {
	//	fmt.Printf("Sentry initialization failed: %v\n", err)
	//}
	// 初始化RocketMQ配置
	//pConfig = rocketmq.ProducerConfig{ClientConfig: rocketmq.ClientConfig{
	//	LogC: &rocketmq.LogConfig{
	//		Path:     "example",
	//		FileSize: 64 * 1 << 10,
	//		FileNum:  1,
	//		Level:    rocketmq.LogLevelDebug,
	//	},
	//}}
	//if err != nil {
	//	panic(err.Error())
	//}
	db.InitDB()
	mem.InitMem()
	// 初始化DAG图
}

func main() {
	initLocal()
	runtime.GOMAXPROCS(runtime.NumCPU())
	schedulerType := os.Args[1]
	switch schedulerType {
	case "dag":
		scheduler.Init()
		scheduler.SetNodeType(dal.NODETYPE_DAG)
		go scheduler.HeartBeat() // slave node
		for true {
			// 只启一个调度器 调度所有分区
			stable.SchedulerLogger.Info("hello dag scheduler ...")
			partitions, err := scheduler.GetMyParitions()
			fmt.Println("my partitions is:", partitions)
			if err != nil {
				stable.SchedulerLogger.Error("GetMyParitions Error!", zap.Error(err))
			}
			//var paritions []int
			//for i := 0; i <= 127; i++ {
			//	paritions = append(paritions, i)
			//}
			if len(partitions) != 0 {
				scheduler.DagScheduler(partitions)
			}
			time.Sleep(500000 * time.Microsecond)
			stable.SchedulerLogger.Info("bye dag scheduler ...")
		}
	case "delay":
		for true {
			stable.SchedulerLogger.Info("hello delay scheduler ...")
			scheduler.DelayScheduler()
			time.Sleep(500000 * time.Microsecond)
			stable.SchedulerLogger.Info("bye delay scheduler ...")
		}
	case "align":
		for true {
			stable.SchedulerLogger.Info("hello align scheduler ...")
			var paritions []int
			for i := 0; i <= 127; i++ {
				paritions = append(paritions, i)
			}
			scheduler.FlowModify(paritions)
			time.Sleep(500000 * time.Microsecond)
			stable.SchedulerLogger.Info("bye align scheduler ...")
		}
	default:
		panic("unknow scheduler type.")
	}
}
