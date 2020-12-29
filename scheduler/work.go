package scheduler

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/stable"
	"go.uber.org/zap"
	"time"
)

func WorkRecover(partitions []int) {
	/*
			1. 认为一个work成功后 和它同类型的任务 在它之前10分钟处于 已下发push &  正在执行running 的任务为卡着的任务
		    2. 对这些任务重置进行重新调度
	*/
	// TODO: 分 partition recover

	// recover
	for _, work := range mem.GlobalTermiteConfig.WorkConfigMap {
		// 获取最近执行成功的work update time
		lastSucWork, err := db.GetLastSuccessWork(work.Key)
		if err != nil {
			stable.SchedulerLogger.Error("GetLastSuccessWork error", zap.Error(err))
			continue
		}

		//  update_time - 20min 之前处于的 push/running状态的认为是IO丢失 进行重置

		etime := lastSucWork.UpdatedAt.Add(-2 * time.Duration(time.Hour))
		err = db.ResetLostWorks(work.Key, etime)
		if err != nil {
			stable.SchedulerLogger.Error("ResetLostWorks error", zap.Error(err))
			continue
		}
	}

	return
}

//func WorkTimeoutRecover(partitions []int) {}
