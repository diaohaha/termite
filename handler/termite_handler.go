package handler

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/dal/prometheus"
	proto "github.com/diaohaha/termite/proto"
	"github.com/diaohaha/termite/stable"
	"context"
	"log"
	"strconv"
)

type TermiteHandler struct{}

func (t *TermiteHandler) AddWorkFlow(ctx context.Context, req *proto.AddWorkFlowRequest, rsp *proto.AddWorkFlowResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "AddWorkFlow").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	// TODO: 事务
	err := db.AddWorkFlow(req.WorkflowKey, req.Project, req.Cid)
	if err != nil {
		stable.CaptureError(err, "AddWorkFlow", "AddWorkFlow", map[string]string{}, map[string]string{
			"method":  "AddWorkFlow",
			"cid":     req.Cid,
			"vflow":   req.WorkflowKey,
			"project": req.Project,
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
		return nil
	} else {
		vworks := mem.GlobalTermiteConfig.FlowConfigMap[req.WorkflowKey].Config.Works
		for _, vwork := range vworks {
			err = db.AddWork(req.WorkflowKey, vwork, req.Project, req.Cid, "")
			if err != nil {
				stable.CaptureError(err, "AddWorkFlow", "AddWorkFlow", map[string]string{}, map[string]string{
					"method":  "AddWorkFlow",
					"cid":     req.Cid,
					"vflow":   req.WorkflowKey,
					"project": req.Project,
				})
				rsp.Code = -1
				rsp.Message = string(err.Error())
				return nil
			}
		}
	}
	return nil
}

func (t *TermiteHandler) FinishWorkFlow(ctx context.Context, req *proto.FinishWorkFlowRequest, rsp *proto.FinishWorkFlowResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "FinishWorkFlow").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	err := db.SetWorkFlowFinish(req.FlowId)
	if err != nil {
		stable.CaptureError(err, "FinishWorkFlow", "FinishWorkFlow", map[string]string{}, map[string]string{
			"method":  "FinishWorkFlow",
			"flow_id": strconv.FormatInt(req.FlowId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) RecoverWorkFlow(ctx context.Context, req *proto.RecoverWorkFlowRequest, rsp *proto.RecoverWorkFlowResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "RecoverWorkFlow").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	err := db.RecoverWorkFlow(req.FlowId)
	if err != nil {
		stable.CaptureError(err, "RecoverWorkFlow", "RecoverWorkFlow", map[string]string{}, map[string]string{
			"method":  "RecoverWorkFlow",
			"flow_id": strconv.FormatInt(req.FlowId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) QueryWork(ctx context.Context, req *proto.QueryWorkRequest, rsp *proto.QueryWorkResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "QueryWork").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	works, err := db.QueryWork(req.Cid, req.WorkflowKey, req.WorkKey)
	if err != nil {
		stable.CaptureError(err, "QueryWorkFlow", "QueryWorkFlow", map[string]string{}, map[string]string{
			"method": "QueryWorkFlow",
			"cid":    req.Cid,
			"vflow":  req.WorkflowKey,
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	var workResults = []*proto.WorkResult{}
	for _, work := range works {
		workResults = append(workResults, &proto.WorkResult{
			Cid:       work.Cid,
			WorkId:    work.Id,
			WorkState: int64(work.Vstate),
		})
	}
	rsp.Works = workResults
	return nil
}

func (t *TermiteHandler) QueryWorkFlow(ctx context.Context, req *proto.QueryWorkFlowRequest, rsp *proto.QueryWorkFlowResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "QueryWorkFlow").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	flowId, err := db.QueryWorkFlow(req.Cid, req.WorkflowKey)
	if err != nil {
		stable.CaptureError(err, "QueryWorkFlow", "QueryWorkFlow", map[string]string{}, map[string]string{
			"method": "QueryWorkFlow",
			"cid":    req.Cid,
			"vflow":  req.WorkflowKey,
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	rsp.FlowId = flowId
	return nil
}

func (t *TermiteHandler) WorkStart(ctx context.Context, req *proto.WorkStartRequest, rsp *proto.WorkStartResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "WorkStart").Trace()
	log.Println("work: ", req.WorkId, " start.")
	rsp.Code = 1
	rsp.Message = "success"
	err := db.SetWorkRunning(req.WorkId)
	log.Println("work: ", req.WorkId, " start1.")
	log.Println(err)
	if err != nil {
		stable.CaptureError(err, "WorkStart", "WorkStart", map[string]string{}, map[string]string{
			"method":  "WorkStart",
			"work_id": strconv.FormatInt(req.WorkId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) WorkFinish(ctx context.Context, req *proto.WorkFinishRequest, rsp *proto.WorkFinishResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "WorkFinish").Trace()
	log.Println("work: ", req.WorkId, " finish.")
	rsp.Code = 1
	rsp.Message = "success"
	err := db.SetWorkFinish(req.WorkId, req.Result, req.Output)
	if err != nil {
		stable.CaptureError(err, "WorkFinish", "WorkFinish", map[string]string{}, map[string]string{
			"method":  "WorkFinish",
			"work_id": strconv.FormatInt(req.WorkId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) WorkError(ctx context.Context, req *proto.WorkErrorRequest, rsp *proto.WorkErrorResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "WorkError").Trace()
	log.Println("work: ", req.WorkId, " error.")
	rsp.Code = 1
	rsp.Message = "success"
	err := db.SetWorkError(req.WorkId, req.Error)
	if err != nil {
		stable.CaptureError(err, "WorkError", "WorkError", map[string]string{}, map[string]string{
			"method":  "WorkError",
			"work_id": strconv.FormatInt(req.WorkId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) WorkDelay(ctx context.Context, req *proto.WorkDelayRequest, rsp *proto.WorkDelayResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "WorkDelay").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	err := db.SetWorkDelay(req.WorkId, req.DelaySeconds)
	if err != nil {
		stable.CaptureError(err, "WorkDelay", "WorkDelay", map[string]string{}, map[string]string{
			"method":  "WorkDelay",
			"work_id": strconv.FormatInt(req.WorkId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}

func (t *TermiteHandler) GetWorkFlowContext(ctx context.Context, req *proto.GetWorkFlowContextRequest, rsp *proto.GetWorkFlowContextResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "GetWorkFlowContext").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	icontext, err := db.GetFlowContext(req.FlowId)
	if err != nil {
		stable.CaptureError(err, "GetWorkFlowContext", "GetWorkFlowContext", map[string]string{}, map[string]string{
			"method":  "GetWorkFlowContext",
			"flow_id": strconv.FormatInt(req.FlowId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	rsp.Context = icontext
	return nil
}

func (t *TermiteHandler) SetWorkFlowContext(ctx context.Context, req *proto.SetWorkFlowContextRequest, rsp *proto.SetWorkFlowContextResponse) error {
	defer prometheus.IOTimeTrace(prometheus.TracerMethod_Rpc, "SetWorkFlowContext").Trace()
	rsp.Code = 1
	rsp.Message = "success"
	err := db.AddFlowContext(req.FlowId, req.Context)
	if err != nil {
		stable.CaptureError(err, "SetWorkFlowContext", "SetWorkFlowContext", map[string]string{}, map[string]string{
			"method":  "SetWorkFlowContext",
			"flow_id": strconv.FormatInt(req.FlowId, 10),
		})
		rsp.Code = -1
		rsp.Message = string(err.Error())
	}
	return nil
}
