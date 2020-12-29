package db

import (
	"github.com/diaohaha/termite/dal"
	"fmt"
)

// 统计相关接口

func GetFlowConfigNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var flowConfig dal.TermiteFlowConfig
	sql := fmt.Sprintf("select count(1) as num from %s", flowConfig.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

func GetWorkConfigNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var workConfig dal.TermiteWorkConfig
	sql := fmt.Sprintf("select count(1) as num from %s", workConfig.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

func GetActiveFlowConfigNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var flow dal.TermiteFlow
	sql := fmt.Sprintf("select count(distinct vflow) as num from %s", flow.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

func GetActiveWorkConfigNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var work dal.TermiteWork
	sql := fmt.Sprintf("select count(distinct vwork) as num from %s", work.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

func GetFlowInstanceNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var flow dal.TermiteFlow
	sql := fmt.Sprintf("select count(1) as num from %s", flow.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

func GetWorkInstanceNum() (num int, err error) {
	var numStru struct {
		Num int `gorm:"column:num"`
	}
	var work dal.TermiteWork
	sql := fmt.Sprintf("select count(1) as num from %s", work.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	num = numStru.Num
	return
}

type NumStateDetail struct {
	Num   int   `gorm:"column:num"`
	State int64 `gorm:"column:state"`
}

func GetFlowInstanceNumDetail() (numStru []NumStateDetail, err error) {
	var flow dal.TermiteFlow
	sql := fmt.Sprintf("select count(1) as num, vstate as state from %s group by state", flow.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	return
}

type NumStateDetailByFlow struct {
	Num   int    `gorm:"column:num"`
	Vflow string `gorm:"column:vflow"`
}

func GetFlowInstanceNumDetailByFlow() (numStru []NumStateDetailByFlow, err error) {
	var flow dal.TermiteFlow
	sql := fmt.Sprintf("select count(1) as num, vflow from %s group by vflow", flow.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	return
}

func GetWorkInstanceNumDetail() (numStru []NumStateDetail, err error) {
	var work dal.TermiteWork
	sql := fmt.Sprintf("select count(1) as num, vstate as state from %s group by state", work.TableName())
	err = RODB.Raw(sql).Scan(&numStru).Error
	return
}
