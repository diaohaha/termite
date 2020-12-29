package db

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"github.com/diaohaha/termite/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"
)

var TerminteStateFsm utils.StateFsm
var TerminteFlowStateFsm utils.StateFsm

func init() {
	// step1: 初始化状态机器配置
	TerminteStateFsm.ConfigParse(utils.ExchangeUnitList{
		utils.ExchangeUnit{dal.TermiteWorkStateInit, dal.TermiteWorkStatePush},
		utils.ExchangeUnit{dal.TermiteWorkStatePush, dal.TermiteWorkStateRunning},
		utils.ExchangeUnit{dal.TermiteWorkStatePush, dal.TermiteWorkStateTimeout},
		utils.ExchangeUnit{dal.TermiteWorkStateRunning, dal.TermiteWorkStateSuccess},
		utils.ExchangeUnit{dal.TermiteWorkStateRunning, dal.TermiteWorkStateFailed},
		utils.ExchangeUnit{dal.TermiteWorkStateRunning, dal.TermiteWorkStateTimeout},
		utils.ExchangeUnit{dal.TermiteWorkStateRunning, dal.TermiteWorkStateDelay},
		utils.ExchangeUnit{dal.TermiteWorkStateRunning, dal.TermiteWorkStateError},
		utils.ExchangeUnit{dal.TermiteWorkStateDelay, dal.TermiteWorkStatePush},
	})
	TerminteFlowStateFsm.ConfigParse(utils.ExchangeUnitList{
		utils.ExchangeUnit{dal.TermiteWorkFlowStateInit, dal.TermiteWorkFlowStateRunning},
		utils.ExchangeUnit{dal.TermiteWorkFlowStateRunning, dal.TermiteWorkFlowStateFinish},
		utils.ExchangeUnit{dal.TermiteWorkFlowStateRunning, dal.TermiteWorkFlowStateTimeout},
	})
}

func AddWorkFlow(vflow string, project string, cid string) error {
	var flow = dal.TermiteFlow{
		Vflow:     vflow,
		Project:   project,
		Cid:       cid,
		Vstate:    dal.TermiteWorkFlowStateInit,
		Partition: rand.Intn(128),
	}
	err := WTDB.Create(&flow).Error
	return err
}

func AddWork(vflow string, vwork string, project string, cid string, env string) error {
	var work = dal.TermiteWork{
		Vstate:  dal.TermiteWorkStateInit,
		Vflow:   vflow,
		Vwork:   vwork,
		Project: project,
		Cid:     cid,
		Env:     env,
	}
	err := WTDB.Create(&work).Error
	return err
}

func SetWorkPush(work_id int64) error {
	var twork dal.TermiteWork
	err := WTDB.Where("id = ?", work_id).First(&twork).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	if err != nil {
		return err
	}
	legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStatePush)
	if legal {
		WTDB.Model(&twork).Update("vstate", dal.TermiteWorkStatePush)
	} else {
		return err
	}
	return nil
}

func SetWorkRunning(work_id int64) error {
	var twork dal.TermiteWork
	err := WTDB.Where("id = ?", work_id).First(&twork).Error
	if err != nil {
		return err
	}
	legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateRunning)
	if legal {
		WTDB.Model(dal.TermiteWork{}).Where("id = ?", twork.Id).Update("vstate", dal.TermiteWorkStateRunning)
	} else {
		return err
	}
	return nil
}

func SetWorkFinish(work_id int64, result string, output string) error {
	var twork dal.TermiteWork
	err := WTDB.Where("id = ?", work_id).First(&twork).Error
	if err != nil {
		return err
	}
	if result == "success" {
		legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateSuccess)
		if legal {
			WTDB.Model(dal.TermiteWork{}).Where("id = ?", twork.Id).Updates(
				dal.TermiteWork{Vstate: dal.TermiteWorkStateSuccess, Output: output})
		} else {
			return err
		}
	} else if result == "fail" {
		legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateFailed)
		if legal {
			WTDB.Model(dal.TermiteWork{}).Where("id = ?", twork.Id).Updates(
				dal.TermiteWork{Vstate: dal.TermiteWorkStateFailed, Output: output})
		} else {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("Unknow Finish Result: %s", result))
	}
	return nil
}

func SetWorkError(work_id int64, error string) error {
	var twork dal.TermiteWork
	err := WTDB.Where("id = ?", work_id).First(&twork).Error
	if err != nil {
		return err
	}
	legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateError)
	if legal {
		WTDB.Model(dal.TermiteWork{}).Where("id = ?", twork.Id).Updates(
			dal.TermiteWork{Vstate: dal.TermiteWorkStateError, Error: error})
		//WTDB.Model(&twork).Update("vstate", dal.TermiteWorkStateError)
		//WTDB.Model(&twork).Update("error", error)
	} else {
		return err
	}
	return nil
}

func SetWorkDelay(work_id int64, delay_seconds int64) (err error) {
	var twork dal.TermiteWork
	err = WTDB.Where("id = ?", work_id).First(&twork).Error
	if err != nil {
		stable.CaptureError(err, "SetWorkDelay", "SetWorkDelay", map[string]string{}, map[string]string{
			"method":  "SetWorkDelay",
			"work_id": strconv.FormatInt(work_id, 10),
		})
		return err
	}
	legal, err := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateDelay)
	if legal {
		delaies := twork.Delaies
		err = WTDB.Model(dal.TermiteWork{}).Where("id = ?", twork.Id).Updates(
			dal.TermiteWork{
				Vstate:  dal.TermiteWorkStateDelay,
				Wakeup:  delay_seconds + time.Now().Unix(),
				Delaies: delaies + 1,
			}).Error
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func SetWorkTimeout(work_id int64) (err error) {
	var twork dal.TermiteWork
	err = WTDB.Where("id = ?", work_id).First(&twork).Error
	if err != nil {
		return err
	}
	legal, _ := TerminteStateFsm.StateCheck(twork.Vstate, dal.TermiteWorkStateTimeout)
	if legal {
		WTDB.Model(&twork).Update("vstate", dal.TermiteWorkStateTimeout)
	}
	return
}

// flow op ...

func SetWorkFlowFinish(flow_id int64) error {
	var tflow dal.TermiteFlow
	err := WTDB.Where("id = ?", flow_id).First(&tflow).Error
	if err != nil {
		return err
	}
	WTDB.Model(&tflow).Update("vstate", dal.TermiteWorkFlowStateFinish)
	err = DeleteFlow(flow_id)
	return err
}

func RecoverWorkFlow(flow_id int64) error {
	var tflow dal.TermiteFlow
	tx := WTDB.Begin()
	err := tx.Where("id = ?", flow_id).First(&tflow).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&tflow).Update("vstate", dal.TermiteWorkFlowStateInit).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(dal.TermiteWork{}).Where(
		"cid = ? and vflow = ?", tflow.Cid, tflow.Vflow,
	).Updates(map[string]interface{}{
		"vstate": dal.TermiteWorkStateInit,
		"output": "",
		"error":  "",
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func RecoverWork(workId int64) error {
	var twork dal.TermiteWork
	tx := WTDB.Begin()
	err := tx.Where("id = ?", workId).First(&twork).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&twork).Update("vstate", dal.TermiteWorkStateInit).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&dal.TermiteFlow{}).Where("cid = ? and vflow = ?", twork.Cid, twork.Vflow).
		Update("vstate", dal.TermiteWorkFlowStateRunning).Error
	tx.Commit()
	return nil
}

func RetryWork(workId int64, retriesLimit int) error {
	// 重试
	var twork dal.TermiteWork
	tx := WTDB.Begin()
	err := tx.Where("id = ?", workId).First(&twork).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if twork.Retries < retriesLimit {
		err = tx.Model(&twork).Update("vstate", dal.TermiteWorkStateInit, "retries", twork.Retries+1).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Model(&dal.TermiteFlow{}).Where("cid = ? and vflow = ?", twork.Cid, twork.Vflow).
			Update("vstate", dal.TermiteWorkFlowStateRunning).Error
		tx.Commit()
	}
	//_ = tx.Close()
	return nil
}

func QueryWork(cid string, vflow string, vwork string) (tworks []dal.TermiteWork, err error) {
	tworks = make([]dal.TermiteWork, 0)
	cur := WTDB.Where("cid = ?", cid)
	if vflow != "" {
		cur = cur.Where("vflow = ?", vflow)
	}
	if vwork != "" {
		cur = cur.Where("vwork = ?", vwork)
	}
	err = cur.Find(&tworks).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if err != nil {
		return
	}
	return
}

func QueryWorkFlow(cid string, vflow string) (flowId int64, err error) {
	var tflow dal.TermiteFlow
	err = WTDB.Where("cid = ? and vflow = ?", cid, vflow).First(&tflow).Error
	if err != nil {
		return
	}
	flowId = tflow.Id
	return
}

func SetWorkFlowTimeout(flow_id int64) error {
	var tflow dal.TermiteFlow
	err := WTDB.Where("id = ?", flow_id).First(&tflow).Error
	if err != nil {
		return err
	}
	WTDB.Model(&tflow).Update("vstate", dal.TermiteWorkFlowStateTimeout)
	return nil
}

func GetTopFlowsByPartitions(partitions []int, length int) []dal.TermiteFlow {
	var tworkflows []dal.TermiteFlow
	WTDB.Where("`tpartition` in (?) and `vstate` in (?)", partitions, []int64{dal.TermiteWorkFlowStateInit, dal.TermiteWorkFlowStateRunning}).
		Order("id desc").Limit(length).Find(&tworkflows)
	return tworkflows
}

func GetFlowsByPartitions(sid int64, partitions []int, vstates []int64, length int) []dal.TermiteFlow {
	var tworkflows []dal.TermiteFlow
	WTDB.Where("`tpartition` in (?) and `vstate` in (?)", partitions, vstates).
		Where("id > ?", sid).
		Order("id asc").Limit(length).Find(&tworkflows)
	return tworkflows
}

func GetFlowsByPartitionsAndVflow(sid int64, partitions []int, vflow string, vstates []int64, length int) []dal.TermiteFlow {
	var tworkflows []dal.TermiteFlow
	WTDB.Where("`tpartition` in (?) and `vstate` in (?) and `vflow` = ?", partitions, vstates, vflow).
		Where("id > ?", sid).
		Order("id asc").Limit(length).Find(&tworkflows)
	return tworkflows
}

func GetFlowsByPartitionsAndVflowOrdered(partitions []int, vflow string, vstates []int64, length int, order string) []dal.TermiteFlow {
	// 排序
	var tworkflows []dal.TermiteFlow
	query := WTDB.Where("`tpartition` in (?) and `vstate` in (?) and `vflow` = ?", partitions, vstates, vflow)
	query = query.
		Order(order).Limit(length).Find(&tworkflows)
	return tworkflows
}

func GetWorksByFlowAndCid(tflow string, cid string) []dal.TermiteWork {
	var tworks []dal.TermiteWork
	WTDB.Where("vflow = ?", tflow).Where("cid = ?", cid).Find(&tworks)
	return tworks
}

func GetTopDelayTasks(sid int64, length int) []dal.TermiteWork {
	var tworks []dal.TermiteWork
	WTDB.Where("vstate = ? and wake_up < ?", dal.TermiteWorkStateDelay, time.Now().Unix()).
		Where("id > ?", sid).
		Order("id asc").Limit(length).Find(&tworks)
	return tworks
}

func GetErrorFlowsByPartitions(partitions []int) ([]dal.TermiteFlow, error) {
	// 取 -1h ~ -30min 的异常任务
	curTime := time.Now()
	var tworkflows []dal.TermiteFlow
	err := WTDB.Where("`tpartition` in (?) and `vstate` in (?)", partitions, []int64{dal.TermiteWorkFlowStateError}).
		Where("create_time > ?", curTime.Add(-time.Minute*60)).
		Where("create_time < ?", curTime.Add(-time.Minute*30)).Find(&tworkflows).Error
	return tworkflows, err
}

//func GetTimeoutFlows(length int) []dal.TermiteFlow {
//	var tflows []dal.TermiteFlow
//	now := time.Now()
//	d, _ := time.ParseDuration("-12h")
//	timeoutTime := now.Add(d)
//	timeStr := timeoutTime.Format("2006-01-02 03:04:05")
//	log.Println(timeStr)
//	WTDB.Where("`vstate` in (?)", []int{dal.TermiteWorkFlowStateInit, dal.TermiteWorkFlowStateRunning}).
//		Where("create_time < ?", timeStr).Order("id desc").Limit(length).Find(&tflows)
//	return tflows
//}

func GetLastSuccessWork(vwork string) (work dal.TermiteWork, err error) {
	// 获取最近成功的task
	err = WTDB.Model(dal.TermiteWork{}).Where(
		"vwork = ? and vstate = ?",
		vwork, dal.TermiteWorkStateSuccess,
	).Find(&work).Order("update_time desc").Error
	return
}

func ResetLostWorks(vwork string, etime time.Time) (err error) {
	// 重置超时的任务
	err = WTDB.Model(dal.TermiteWork{}).Where(
		"vwork = ? and update_time <= ? and vstate in (?)",
		vwork, etime, []int64{dal.TermiteWorkStatePush, dal.TermiteWorkStateRunning},
	).Update("vstate", dal.TermiteWorkStateInit).Error
	return
}

func GetRunningWorkBefore(before time.Time) (works []dal.TermiteWork, err error) {
	err = WTDB.Where("vstate = ? and update_time < ?", dal.TermiteWorkStateRunning, before).Find(&works).Error
	if err != nil {
		return
	}
	return
}
func GetPushWorkBefore(before time.Time) (works []dal.TermiteWork, err error) {
	err = WTDB.Where("vstate = ? and update_time < ?", dal.TermiteWorkStatePush, before).Find(&works).Error
	if err != nil {
		return
	}
	return
}

func GetTimeoutWorks() (works []dal.TermiteWork, err error) {
	err = WTDB.Model(dal.TermiteWork{}).Where("vstate = ?", dal.TermiteWorkStateTimeout).Find(&works).Error
	return
}
