package db

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
)

func GetTermiteWorkConfigs() (ret []dal.TermiteWorkConfig, err error) {
	err = RODB.Find(&ret).Error
	return ret, err
}

func GetTermiteWorkConfig(workKey string) (ret dal.TermiteWorkConfig, err error) {
	err = RODB.Where("work_key = ?", workKey).First(&ret).Error
	return ret, err
}

func GetTermiteFlowConfigs() (ret []dal.TermiteFlowConfig, err error) {
	err = RODB.Find(&ret).Error
	return ret, err
}

func GetTermiteFlowConfig(flowKey string) (ret dal.TermiteFlowConfig, err error) {
	err = RODB.Where("flow_key = ?", flowKey).First(&ret).Error
	return ret, err
}

func UpdateTermiteFlowConfig(config dal.TermiteFlowConfig) error {
	err := WTDB.Save(&config).Error
	return err
}

func UpdateTermiteWorkConfig(config dal.TermiteWorkConfig) error {
	err := WTDB.Save(&config).Error
	return err
}

func DeleteTermiteFlowConfig(flow_key string) {
	WTDB.Delete(dal.TermiteFlowConfig{}, "flow_key = ?", flow_key)
}

func DeleteTermiteWorkConfig(work_key string) {
	WTDB.Delete(dal.TermiteWorkConfig{}, "work_key = ?", work_key)
}

func QueryTermiteFlowConfigs(vflow string, search string, pageIndex int, pageSize int) (flowConfigs []dal.TermiteFlowConfig, count int, err error) {
	// 查询工作流配置
	query := RODB.Model(dal.TermiteFlowConfig{})
	if vflow != "" {
		query = query.Where("flow_key = ?", vflow)
	} else {
		if search != "" {
			query = query.Where("flow_key like ?", "%"+search+"%")
		}
	}
	err = query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&flowConfigs).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteFlowConfigs Error")
	}
	err = query.Count(&count).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteFlowConfigs Error")
	}
	return
}

func QueryTermiteWorkConfigs(vwork string, search string, pageIndex int, pageSize int) (workConfigs []dal.TermiteWorkConfig, count int, err error) {
	// 查询工作配置
	query := RODB.Model(dal.TermiteWorkConfig{})
	if vwork != "" {
		query = query.Where("work_key = ?", vwork)
	} else {
		if search != "" {
			query = query.Where("work_key like ?", "%"+search+"%")
		}
	}
	err = query.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&workConfigs).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteWorkConfigs Error")
	}
	err = query.Count(&count).Error
	if err != nil {
		stable.Logger.Error("QueryTermiteWorkConfigs Error")
	}
	return
}

func UpdateFlowSwitch(vflow string, iswitch int) (err error) {
	// 更新工作流开关
	var flowConfig dal.TermiteFlowConfig
	err = RODB.Model(dal.TermiteFlowConfig{}).Where("flow_key = ?", vflow).Find(&flowConfig).Error
	if err != nil {
		stable.Logger.Error("UpdateFlowSwitch Error")
		return
	}
	flowConfig.Switch = iswitch
	err = WTDB.Save(&flowConfig).Error
	if err != nil {
		stable.Logger.Error("UpdateFlowSwitch Error")
		return
	}
	return
}
