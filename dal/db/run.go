package db

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/redis"
	"encoding/json"
)

func AddFlowContext(flowId int64, context map[string]string) (err error) {
	err = redis.Lock(flowId)
	if err != nil {
		return err
	}
	oldContext, err := GetFlowContext(flowId)
	if err != nil {
		return err
	}
	for newKey, newValue := range context {
		oldContext[newKey] = newValue
	}
	contextStr, err := json.Marshal(oldContext)
	if err != nil {
		return err
	}
	var flow dal.TermiteFlow
	err = WTDB.Where("id = ?", flowId).First(&flow).Error
	if err != nil {
		return err
	}
	//flow.Context = string(contextStr)
	//err = WTDB.Save(&flow).Error
	err = WTDB.Model(dal.TermiteFlow{}).Where("id = ?", flowId).Update("tcontext", string(contextStr)).Error
	_ = redis.UnLock(flowId)
	return err
}

func GetFlowContext(flowId int64) (context map[string]string, err error) {
	context = make(map[string]string, 0)
	var flow dal.TermiteFlow
	err = RODB.Where("id = ?", flowId).First(&flow).Error
	if err != nil {
		return
	}
	if flow.Context == "" {
		return
	}
	err = json.Unmarshal([]byte(flow.Context), &context)
	if err != nil {
		return
	}
	return
}

func GetFlowByFlowKeyAndCid(flowKey string, cid string) (dal.TermiteFlow, error) {
	var flow dal.TermiteFlow
	err := WTDB.Where("vflow = ? and cid = ?", flowKey, cid).First(&flow).Error
	if err != nil {
		return flow, err
	}
	return flow, nil
}

func UpdateFlowState(flowId int64, state int) error {
	var flow dal.TermiteFlow
	err := WTDB.Where("id = ?", flowId).First(&flow).Error
	if err != nil {
		return err
	}
	if flow.Vstate == int32(state) {
		return nil
	}
	//flow.Vstate = int32(state)
	//err = WTDB.Save(flow).Error
	err = WTDB.Model(dal.TermiteFlow{}).Where("id = ?", flowId).Update("vstate", int32(state)).Error
	return err
}

func DeleteFlow(flowId int64) error {
	var flow dal.TermiteFlow
	err := WTDB.Where("id = ?", flowId).First(&flow).Error
	if err != nil {
		return err
	}
	tx := WTDB.Begin()

	err = tx.Delete(&flow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("cid = ? and vflow = ?", flow.Cid, flow.Vflow).Delete(dal.TermiteWork{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
