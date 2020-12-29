package main

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/scheduler"
	"github.com/diaohaha/termite/stable"
	"github.com/diaohaha/termite/stable/alert"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"os"
)

func initLocal() {
	// sentry config
	//err := sentry.Init(sentry.ClientOptions{
	//	Dsn: dal.Env.Sentry_Dsn,
	//})
	//if err != nil {
	//	fmt.Printf("Sentry initialization failed: %v\n", err)
	//}
	db.InitDB()
	mem.InitMem()
	scheduler.Init()
}

func main() {
	/*
	   1. 定时检查超时任务
	*/
	sigs := make(chan os.Signal, 1)
	bIsRunningWorkTimeout := false
	//bIsRunningFlowModify := false
	//bIsRunningFlowRecover := false
	//bIsRunningWorkRecover := false
	bIsRunningSchedulerMaster := false
	bIsRunningWorkRetryer := false
	c := cron.New()
	initLocal()

	_ = c.AddFunc("@every 300s", func() {
		// 监控检查
		alert.AlertCheck()
	})

	_ = c.AddFunc("@every 60s", func() {
		// 检查任务超时
		stable.CronLogger.Info("hello. work timeout checker")
		if bIsRunningWorkTimeout {
			stable.CronLogger.Info("work timeout checker still running")
			return
		}
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			bIsRunningWorkTimeout = false
			if err := recover(); err != nil {
				stable.SchedulerLogger.Error("WorkTimeoutCheck panic", zap.Error(err.(error))) // 这里的err其实就是panic传入的内容
				stable.CaptureError(err.(error), "WorkTimeoutRetry", "WorkTimeoutRetry", map[string]string{}, map[string]string{
					"method": "workExecTimeoutCheck",
					"type":   "panic",
				})
				return
			}
		}()
		bIsRunningWorkTimeout = true
		scheduler.WorkTimeoutCheck()
		stable.CronLogger.Info("byebye. work timeout checker")
	})

	_ = c.AddFunc("@every 10s", func() {
		// 检查任务超时
		stable.CronLogger.Info("hello. work retry checker")
		if bIsRunningWorkRetryer {
			stable.CronLogger.Info("retry checker still running")
			return
		}
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			bIsRunningWorkRetryer = false
			if err := recover(); err != nil {
				stable.SchedulerLogger.Error("WorkTimeoutRetry panic", zap.Error(err.(error))) // 这里的err其实就是panic传入的内容
				stable.CaptureError(err.(error), "WorkTimeoutRetry", "WorkTimeoutRetry", map[string]string{}, map[string]string{
					"method": "WorkTimeoutRetry",
					"type":   "panic",
				})
				return
			}
		}()
		bIsRunningWorkRetryer = true
		scheduler.WorkTimeoutRetry()
		stable.CronLogger.Info("byebye. work retry checker")
	})

	//_ = c.AddFunc("@every 60s", func() {
	//	// 根据work状态修改flow状态
	//	// TODO: 定位不能recover问题
	//	stable.CronLogger.Info("hello. flow modify")
	//	//if bIsRunningFlowModify {
	//	//	return
	//	//}
	//	bIsRunningFlowModify = true
	//	var paritions []int
	//	for i := 0; i <= 127; i++ {
	//		paritions = append(paritions, i)
	//	}
	//	scheduler.FlowModify(paritions)
	//	bIsRunningFlowModify = false
	//	stable.CronLogger.Info("byebye. flow modify")
	//})

	//_ = c.AddFunc("0 */30 * * * ?", func() {
	//	// Exception任务进行一次自动修复 每隔30min
	//	stable.CronLogger.Info("hello. flow recover")
	//	if bIsRunningFlowRecover {
	//		return
	//	}
	//	bIsRunningFlowModify = true
	//	var paritions []int
	//	for i := 0; i <= 127; i++ {
	//		paritions = append(paritions, i)
	//	}
	//	scheduler.FlowRecover(paritions)
	//	bIsRunningFlowRecover = false
	//	stable.CronLogger.Info("byebye. flow recover")
	//})

	//_ = c.AddFunc("0 */60 * * * ?", func() {
	//	// pushed runninng 任务recover
	//	stable.CronLogger.Info("hello. work recover")
	//	if bIsRunningWorkRecover {
	//		return
	//	}
	//	bIsRunningWorkRecover = true
	//	var paritions []int
	//	for i := 0; i <= 127; i++ {
	//		paritions = append(paritions, i)
	//	}
	//	scheduler.WorkRecover(paritions)
	//	bIsRunningWorkRecover = false
	//	stable.CronLogger.Info("byebye. work recover")
	//})

	_ = c.AddFunc("*/5 * * * * ?", func() {
		// schduler master
		stable.CronLogger.Info("hello. rebalancer")
		//if bIsRunningSchedulerMaster {
		//	stable.CronLogger.Info("rebalancer still running")
		//	return
		//}
		defer func() { bIsRunningSchedulerMaster = false }()
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				stable.SchedulerLogger.Error("SchedulerMaster panic", zap.Error(err.(error))) // 这里的err其实就是panic传入的内容
				stable.CaptureError(err.(error), "SchedulerMaster", "SchedulerMaster", map[string]string{}, map[string]string{
					"method": "SchedulerMaster",
					"type":   "panic",
				})
				return
			}
		}()
		bIsRunningSchedulerMaster = true
		scheduler.Rebalance()
		scheduler.DeleteExpireNodes()
		stable.CronLogger.Info("byebye. rebalancer")
	})

	c.Start()
	<-sigs
}
