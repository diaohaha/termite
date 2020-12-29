package router

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/middleware"
	"github.com/labstack/echo"
)

type FlowGlobalInfo struct {
	FlowConfigNum        int           `json:"flow_config_num"`
	ActiveFlowNum        int           `json:"active_flow_num"`
	FlowInstanceNum      int           `json:"flow_instance_num"`
	FlowNumDetailByState []interface{} `json:"flow_num_detail_by_state"`
	FlowNumDetailByFlow  []interface{} `json:"flow_num_detail_by_flow"`
}

func vInfoFlowCount(ctx echo.Context) (err error) {
	// 工作流全局统计
	var data FlowGlobalInfo
	data.FlowConfigNum, err = db.GetFlowConfigNum()
	data.ActiveFlowNum, err = db.GetActiveFlowConfigNum()
	data.FlowInstanceNum, err = db.GetFlowInstanceNum()
	numDetail, err := db.GetFlowInstanceNumDetail()
	for _, numDetailItem := range numDetail {
		data.FlowNumDetailByState = append(data.FlowNumDetailByState, map[string]interface{}{
			"state": dal.GetWorkFlowStateZh(int32(numDetailItem.State)),
			"num":   numDetailItem.Num,
		})
	}
	numDetailByFlow, err := db.GetFlowInstanceNumDetailByFlow()
	for _, numDetailByFlowItem := range numDetailByFlow {
		data.FlowNumDetailByFlow = append(data.FlowNumDetailByFlow, map[string]interface{}{
			"vflow": numDetailByFlowItem.Vflow,
			"num":   numDetailByFlowItem.Num,
		})
	}
	return middleware.ApiResult(ctx, err, "", "", data)
}

type WorkGlobalInfo struct {
	WorkConfigNum   int           `json:"work_config_num"`
	ActiveWorkNum   int           `json:"active_work_num"`
	WorkInstanceNum int           `json:"work_instance_num"`
	WorkNumDetail   []interface{} `json:"work_num_detail"`
}

func vInfoWorkCount(ctx echo.Context) (err error) {
	// 工作全局统计
	var data WorkGlobalInfo
	data.WorkConfigNum, err = db.GetWorkConfigNum()
	data.ActiveWorkNum, err = db.GetActiveWorkConfigNum()
	data.WorkInstanceNum, err = db.GetWorkInstanceNum()
	numDetail, err := db.GetWorkInstanceNumDetail()
	for _, numDetailItem := range numDetail {
		data.WorkNumDetail = append(data.WorkNumDetail, map[string]interface{}{
			"state": dal.GetWorkStateZh(int32(numDetailItem.State)),
			"num":   numDetailItem.Num,
		})
	}
	return middleware.ApiResult(ctx, err, "", "", data)
}
