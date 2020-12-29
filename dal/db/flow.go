package db

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"math/rand"
)

func QueryTermiteFlowInstances(cid string, vflow string, vstates []int, pageIndex int, pageSize int) (flows []dal.TermiteFlow, count int, err error) {
	// 查询工作流实例
	query := RODB.Model(dal.TermiteFlow{})
	if cid != "" {
		query = query.Where("cid = ?", cid)
	}
	if vflow != "" {
		query = query.Where("vflow = ?", vflow)
	}
	if len(vstates) != 0 {
		query = query.Where("vstate in (?)", vstates)
	}
	err = query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&flows).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteFlowInstances Error")
	}
	err = query.Count(&count).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteFlowInstances Error")
	}
	return
}

func QueryTermiteWorkInstances(cid string, vflow string, vwork string, pageIndex int, pageSize int) (works []dal.TermiteWork, count int, err error) {
	// 查询工作流实例
	query := RODB.Model(dal.TermiteWork{})
	if cid != "" {
		query = query.Where("cid = ?", cid)
	}
	if vflow != "" {
		query = query.Where("vflow = ?", vflow)
	}
	if vwork != "" {
		query = query.Where("vwork = ?", vwork)
	}
	err = query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&works).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteWorkInstances Error")
	}
	err = query.Count(&count).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteWorkInstances Error")
	}
	return
}

func AddWorkFlowIns(vflow string, vworks []string, cid string) (err error) {
	// 添加工作流实例

	tx := WTDB.Begin()

	var flow = dal.TermiteFlow{
		Vflow:     vflow,
		Cid:       cid,
		Vstate:    dal.TermiteWorkFlowStateInit,
		Partition: rand.Intn(128),
	}
	err = tx.Create(&flow).Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, vwork := range vworks {
		var work = dal.TermiteWork{
			Vstate: dal.TermiteWorkStateInit,
			Vflow:  vflow,
			Vwork:  vwork,
			Cid:    cid,
			Env:    "",
		}
		err = tx.Create(&work).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}

func GetFlowCountByStates(vflow string, states []int64) (count int, err error) {
	// 根据状态获取工作流实例数量
	err = RODB.Model(dal.TermiteFlow{}).Where("vflow = ? and vstate in (?)", vflow, states).Count(&count).Error
	return
}
