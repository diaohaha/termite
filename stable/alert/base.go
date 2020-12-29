package alert

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/stable"
	"strconv"
	"time"
)

// check exception level

const (
	LEVEL_CRITICAL         = "critical"
	LEVEL_IMPORTANT        = "important"
	LEVEL_WARNING          = "warning"
	WORKERROR_NUM_THREHOLD = 50
	WORKERROR_SECONDS      = 300
	WORKFAIL_NUM_THREHOLD  = 50
	WORKFAIL_SECONDS       = 300
)

type TermiteAlertError struct {
	Message   string            `json:"message"`
	Item      string            `json:"item"`
	Level     string            `json:"level"`
	StartTime time.Time         `json:"start_time"`
	Info      map[string]string `json:"info"`
	Url       string            `json:"url"`
	Detail    string            `json:"detail"`
}

func AlertCheck() {
	for _, workConfig := range mem.GlobalTermiteConfig.WorkConfigMap {
		if a, ok := CheckReceiveLargeWorkFail(workConfig.Key); ok {
			_ = SendDingMsg(a)
		}
		if a, ok := CheckReceiveLargeWorkError(workConfig.Key); ok {
			_ = SendDingMsg(a)
		}
	}
}

func CheckReceiveLargeWorkError(vwork string) (talert TermiteAlertError, ok bool) {
	// Receive large error in a short time
	var count int
	var checkTime time.Time
	checkTime = time.Now()
	query := db.RODB.Model(dal.TermiteWork{}).Where(
		"vwork = ? and update_time > ? and vstate = ?",
		vwork, checkTime.Add(-WORKERROR_SECONDS*time.Second), dal.TermiteWorkStateError)
	err := query.Count(&count).Error
	if err != nil {
		stable.CaptureError(err, "CheckReceiveLargeWorkError", "", map[string]string{}, map[string]string{
			"method": "CheckReceiveLargeWorkError",
		})
	}
	if count > WORKERROR_NUM_THREHOLD {
		ok = true
		var tmpWork dal.TermiteWork
		err = db.RODB.Model(dal.TermiteWork{}).Where("vwork = ? and vstate = ?", vwork,
			dal.TermiteWorkStateError).Order("id desc").Limit(1).Find(&tmpWork).Error
		if err != nil {
			stable.CaptureError(err, "CheckReceiveLargeWorkError", "", map[string]string{}, map[string]string{
				"method": "CheckReceiveLargeWorkError",
			})
		}
		talert = TermiteAlertError{
			Message:   "receive too many work error request!",
			Item:      "任务异常频率",
			Level:     LEVEL_CRITICAL,
			StartTime: checkTime,
			Url:       "http://172.17.2.77:10023/admin#/iworks/" + vwork,
			Detail:    tmpWork.Error,
			Info: map[string]string{
				"任务名称":  mem.GlobalTermiteConfig.WorkConfigMap[tmpWork.Vwork].Name,
				"任务ID":  tmpWork.Vwork,
				"所属工作流": mem.GlobalTermiteConfig.FlowConfigMap[tmpWork.Vflow].Name,
				"工作流ID": tmpWork.Vflow,
				"检查时间":  checkTime.Add(-WORKFAIL_SECONDS*time.Second).Format("03:04:05") + "-" + checkTime.Format("03:04:05"),
				"异常任务数": strconv.Itoa(count),
			},
		}
	}
	return
}

func CheckReceiveLargeWorkFail(vwork string) (talert TermiteAlertError, ok bool) {
	// Receive large fail in a short time
	var count int
	var checkTime time.Time
	checkTime = time.Now()
	query := db.RODB.Model(dal.TermiteWork{}).Where(
		"vwork = ? and update_time > ? and vstate = ?",
		vwork, checkTime.Add(-WORKFAIL_SECONDS*time.Second), dal.TermiteWorkStateFailed)
	err := query.Count(&count).Error
	if err != nil {
		stable.CaptureError(err, "CheckReceiveLargeWorkFail", "", map[string]string{}, map[string]string{
			"method": "CheckReceiveLargeWorkFail",
		})
	}
	if count > WORKFAIL_NUM_THREHOLD {
		ok = true
		var tmpWork dal.TermiteWork
		err = db.RODB.Model(dal.TermiteWork{}).Where("vwork = ? and vstate = ?", vwork,
			dal.TermiteWorkStateFailed).Order("id desc").Limit(1).Find(&tmpWork).Error
		if err != nil {
			stable.CaptureError(err, "CheckReceiveLargeWorkFail", "", map[string]string{}, map[string]string{
				"method": "CheckReceiveLargeWorkFail",
			})
		}
		talert = TermiteAlertError{
			Message:   "receive too many work fail request!",
			Item:      "任务失败频率",
			Level:     LEVEL_IMPORTANT,
			StartTime: checkTime,
			Url:       "http://172.17.2.77:10023/admin#/iworks/" + vwork,
			Detail:    tmpWork.Output,
			Info: map[string]string{
				"任务名称":  mem.GlobalTermiteConfig.WorkConfigMap[tmpWork.Vwork].Name,
				"任务ID":  tmpWork.Vwork,
				"所属工作流": mem.GlobalTermiteConfig.FlowConfigMap[tmpWork.Vflow].Name,
				"工作流ID": tmpWork.Vflow,
				"检查时间":  checkTime.Add(-WORKFAIL_SECONDS*time.Second).Format("03:04:05") + "-" + checkTime.Format("03:04:05"),
				"任务失败数": strconv.Itoa(count),
			},
		}
	}
	return
}
