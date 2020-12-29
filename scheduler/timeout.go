package scheduler

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/stable"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func WorkTimeoutCheck() {
	workExecTimeoutCheck()
	workPushTimeoutCheck()
}

func workPushTimeoutCheck() {
	// 任务IO丢失 - 超时
	m, _ := time.ParseDuration("-1h") // protect
	fmt.Println("任务下发超时检查...")
	lastSuccessWorkMap := make(map[string]time.Time, 0)
	for _, workConfig := range mem.GlobalTermiteConfig.WorkConfigMap {
		work, err := db.GetLastSuccessWork(workConfig.Key)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
			} else {
				stable.CaptureError(err, "workPushTimeoutCheck", "workPushTimeoutCheck", map[string]string{}, map[string]string{
					"method": "GetLastSuccessWork",
				})
			}
		} else {
			lastSuccessWorkMap[workConfig.Key] = work.UpdatedAt
		}
	}
	works, err := db.GetPushWorkBefore(time.Now().Add(m))
	if err != nil {
		stable.SchedulerLogger.Error("GetPushWorkBefore", zap.Error(err))
		stable.CaptureError(err, "workPushTimeoutCheck", "workPushTimeoutCheck", map[string]string{}, map[string]string{
			"method": "workExecTimeoutCheck",
		})
		return
	}
	count := 0
	sum := 0
	for _, work := range works {
		sum = sum + 1
		fmt.Println("work_id", work.Id, "vwork:", work.Vwork, "超时时间: ", mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].PushTimeout, "已经过了 ", (time.Now().Unix() - work.UpdatedAt.Unix()))
		if _, ok := lastSuccessWorkMap[work.Vwork]; !ok {
			// protect 如果无未成功的task 跳过
			fmt.Println("无更早success任务跳过")
			continue
		}
		if mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].PushTimeout == 0 {
			// 0 跳过超时
			fmt.Println("超时设置为0跳过")
			continue
		}
		if mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].PushTimeout < (time.Now().Unix() - work.UpdatedAt.Unix()) {
			// 超时
			if work.UpdatedAt.Before(lastSuccessWorkMap[work.Vwork]) {
				ierr := db.SetWorkTimeout(work.Id)
				if ierr != nil {
					stable.CaptureError(err, "workPushTimeoutCheck", "workPushTimeoutCheck", map[string]string{}, map[string]string{
						"method":  "workPushTimeoutCheck",
						"cid":     work.Cid,
						"vwork":   work.Vwork,
						"vflow":   work.Vflow,
						"work_id": strconv.FormatInt(work.Id, 10),
					})
					stable.SchedulerLogger.Error("SetWorkTimeoutError", zap.Error(ierr))
				} else {
					count = count + 1
				}
			}
		}
	}
	fmt.Println("任务下发超时检查结束，检查任务个数:", sum, "超时任务个数:", count)
	// 超时时间 - 10
}

func workExecTimeoutCheck() {
	// 任务异常中断 - 超时
	m, _ := time.ParseDuration("-5s") // protect
	fmt.Println("任务执行超时检查...")
	works, err := db.GetRunningWorkBefore(time.Now().Add(m))
	if err != nil {
		stable.SchedulerLogger.Error("GetRunningWorkBeforeError", zap.Error(err))
		stable.CaptureError(err, "WorkTimeoutRetry", "WorkTimeoutRetry", map[string]string{}, map[string]string{
			"method": "workExecTimeoutCheck",
		})
		return
	}
	count := 0
	for _, work := range works {
		if mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].ExecTimeout == 0 {
			// 0 跳过超时
			continue
		}
		if mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].ExecTimeout < (time.Now().Unix() - work.UpdatedAt.Unix()) {
			// 超时
			ierr := db.SetWorkTimeout(work.Id)
			if ierr != nil {
				stable.CaptureError(err, "workExecTimeoutCheck", "workExecTimeoutCheck", map[string]string{}, map[string]string{
					"method":  "workExecTimeoutCheck",
					"cid":     work.Cid,
					"vwork":   work.Vwork,
					"vflow":   work.Vflow,
					"work_id": strconv.FormatInt(work.Id, 10),
				})
				stable.SchedulerLogger.Error("SetWorkTimeoutError", zap.Error(ierr))
			} else {
				count = count + 1
			}
		}
	}
	fmt.Println("任务超时检查结束，超时任务个数：", count)
}

func WorkTimeoutRetry() {
	fmt.Println("WorkTimeoutRetry Start...")
	works, err := db.GetTimeoutWorks()
	if err != nil {
		stable.CaptureError(err, "WorkTimeoutRetry", "WorkTimeoutRetry", map[string]string{}, map[string]string{
			"method": "WorkTimeoutRetry",
		})
		return
	}
	for _, work := range works {
		fmt.Println(mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].Retries)
		ierr := db.RetryWork(work.Id, mem.GlobalTermiteConfig.WorkConfigMap[work.Vwork].Retries)
		if ierr != nil {
			stable.CaptureError(err, "WorkTimeoutRetry", "WorkTimeoutRetry", map[string]string{}, map[string]string{
				"method":  "WorkTimeoutRetry",
				"cid":     work.Cid,
				"vwork":   work.Vwork,
				"vflow":   work.Vflow,
				"work_id": strconv.FormatInt(work.Id, 10),
			})
			stable.SchedulerLogger.Error("workTimeoutRetry", zap.Error(ierr))
		}
	}
	fmt.Println("WorkTimeoutRetry End...")
}
