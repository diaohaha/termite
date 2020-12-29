package router

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/middleware"
	"github.com/diaohaha/termite/utils"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func vQueryFlowConfig(ctx echo.Context) (err error) {
	// parse params
	var flowKey string
	flowKey, err = middleware.GetStringFromCtx(ctx, "flow_key")

	flowConfigMap := mem.GetFlowConfigMap()
	var flowConfigs []mem.CacheTermiteFlowConfig
	if flowKey == "" {
		for _, flowConfig := range flowConfigMap {
			flowConfigs = append(flowConfigs, flowConfig)
		}
	} else {
		for flowName, flowConfig := range flowConfigMap {
			if strings.Contains(flowName, flowKey) {
				flowConfigs = append(flowConfigs, flowConfig)
			}
		}
	}
	// map to list
	_, _ = json.Marshal(flowConfigs)
	//log.Println(string(i))
	if len(flowConfigs) > 0 {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": flowConfigs,
		})
	} else {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": []mem.CacheTermiteFlowConfig{},
		})
	}
}

func vDeleteFlowConfig(ctx echo.Context) error {
	var flowKey string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求参数错误",
		})
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	flowKey, _ = reqDataJs.Get("flow_key").String()
	err = utils.CheckNaminConventions(flowKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "flow_key 命名不规范",
		})
	}
	db.DeleteTermiteFlowConfig(flowKey)
	go mem.ForceRefresh()
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "删除成功",
	})
}

func vConfigCreateFlow(ctx echo.Context) error {
	// parse params
	var flowKey string
	var flowDesc string
	var flowName string
	var config string
	var env string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		flowKey = ""
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	flowKey, err = reqDataJs.Get("flow_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 flow_key 解析失败",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[flowKey]; ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "已经存在",
		})
	}
	flowName, err = reqDataJs.Get("flow_name").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 flow_name 解析失败",
		})
	}
	flowDesc, _ = reqDataJs.Get("flow_desc").String()
	config, err = reqDataJs.Get("config").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "依赖配置解析失败",
		})
	}
	res, err := mem.CheckFlowConfig(config)
	//log.Println(config)
	if !res {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "依赖配置错误: " + string(err.Error()),
		})
	}
	env, _ = reqDataJs.Get("env").String()
	// 检查flowkey是否重复
	var flowConfig dal.TermiteFlowConfig
	flowConfig.FlowKey = flowKey
	flowConfig.FlowName = flowName
	flowConfig.FlowDesc = flowDesc
	flowConfig.FlowConfig = config
	flowConfig.Env = env
	err = db.UpdateTermiteFlowConfig(flowConfig)
	go mem.ForceRefresh()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "创建失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "创建成功",
	})
}

func vConfigCopyFlow(ctx echo.Context) error {
	// parse params
	var srcFlowKey string
	var flowKey string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		flowKey = ""
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	flowKey, err = reqDataJs.Get("flow_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 flow_key 解析失败",
		})
	}
	err = utils.CheckNaminConventions(flowKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "flow_key 命名不规范",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[flowKey]; ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "已经存在",
		})
	}
	srcFlowKey, err = reqDataJs.Get("src_flow_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 src_flow_key 解析失败",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[srcFlowKey]; !ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "src_flow_key不存在",
		})
	}
	dflow, err := db.GetTermiteFlowConfig(srcFlowKey)
	//log.Println("0000 flowConfig", dflow)
	var flowConfig dal.TermiteFlowConfig
	flowConfig.FlowKey = flowKey
	flowConfig.FlowName = dflow.FlowName
	flowConfig.FlowDesc = dflow.FlowDesc
	flowConfig.FlowConfig = dflow.FlowConfig
	flowConfig.Env = dflow.Env
	//log.Println("0000 flowConfig", flowConfig)
	err = db.UpdateTermiteFlowConfig(flowConfig)
	go mem.ForceRefresh()
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "复制失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "复制成功",
	})
}

func vConfigUpdateFlow(ctx echo.Context) error {
	// parse params
	var flowKey string
	var flowDesc string
	var flowName string
	var config string
	var env string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		flowKey = ""
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	flowKey, err = reqDataJs.Get("flow_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 flow_key 解析失败",
		})
	}
	err = utils.CheckNaminConventions(flowKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "flow_key 命名不规范",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.FlowConfigMap[flowKey]; !ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "工作流不存在",
		})
	}
	flowName, err = reqDataJs.Get("flow_name").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数 flow_name 解析失败",
		})
	}
	flowDesc, _ = reqDataJs.Get("flow_desc").String()
	config, err = reqDataJs.Get("config").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "依赖配置解析失败",
		})
	}
	//log.Println(config)
	res, err := mem.CheckFlowConfig(config)
	if !res {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "依赖配置错误:" + string(err.Error()),
		})
	}
	env, _ = reqDataJs.Get("env").String()
	// 检查flowkey是否重复
	flowConifg, err := db.GetTermiteFlowConfig(flowKey)
	flowConifg.Env = env
	flowConifg.FlowConfig = config
	flowConifg.FlowDesc = flowDesc
	flowConifg.FlowName = flowName
	err = db.UpdateTermiteFlowConfig(flowConifg)
	go mem.ForceRefresh()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "更新失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "更新成功",
	})
}

func configGetWorks(ctx echo.Context) error {
	//tc := db.TermiteConfig{}
	//flows := tc.GetFlows(wDB)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"pageTotal": 2,
		"list": []map[string]interface{}{
			{
				"id":   "1",
				"name": "upload",
				"desc": "视频上传",
				"config": map[string]string{
					"template": "h264_0101",
				},
			},
			{
				"id":   "2",
				"name": "get_video_info",
				"desc": "获取视频信息",
				"config": map[string]string{
					"logo_path": "logo/logo_v8.jpg",
				},
			},
			{
				"id":   "2",
				"name": "video_filter1st",
				"desc": "视频第一次过滤",
				"config": map[string]string{
					"logo_path": "logo/logo_v8.jpg",
				},
			},
			{
				"id":   "2",
				"name": "dedup_md5",
				"desc": "md5消重",
			},
			{
				"id":   "2",
				"name": "dedup_all",
				"desc": "策略消重",
			},
			{
				"id":   "2",
				"name": "porn_detect",
				"desc": "色情识别",
			},
			{
				"id":   "2",
				"name": "terrsiom_detect",
				"desc": "暴恐识别",
			},
			{
				"id":   "2",
				"name": "policty",
				"desc": "涉政识别",
			},
			{
				"id":   "2",
				"name": "logo_detect",
				"desc": "水印识别",
			},
			{
				"id":   "2",
				"name": "get_video_frame",
				"desc": "原始视频抽帧",
				"config": map[string]string{
					"logo_path": "logo/logo_v8.jpg",
				},
			},
			{
				"id":   "2",
				"name": "watermark",
				"desc": "水印",
				"config": map[string]string{
					"logo_path": "logo/logo_v8.jpg",
				},
			},
			{
				"id":   "2",
				"name": "watermark",
				"desc": "水印",
				"config": map[string]string{
					"logo_path": "logo/logo_v8.jpg",
				},
			},
		},
	})
}

// Work Config

func vQueryWorkConfig(ctx echo.Context) (err error) {
	var workKey string
	workKey, err = middleware.GetStringFromCtx(ctx, "work_key")

	workConfigMap := mem.GetWorkConfigMap()
	var workConfigs []mem.CacheTermiteWorkConfig
	if workKey == "" {
		for _, workConfig := range workConfigMap {
			workConfigs = append(workConfigs, workConfig)
		}
	} else {
		for workName, workConfig := range workConfigMap {
			if strings.Contains(workName, workKey) {
				workConfigs = append(workConfigs, workConfig)
			}
		}
	}
	// map to list
	if len(workConfigs) > 0 {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": workConfigs,
		})
	} else {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "A0001",
			"data": []mem.CacheTermiteWorkConfig{},
		})
	}
}

func vConfigCreateWork(ctx echo.Context) error {
	var workKey string
	var workDesc string
	var workName string
	var config string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求参数错误",
		})
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	workKey, err = reqDataJs.Get("work_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数work_key解析失败",
		})
	}
	err = utils.CheckNaminConventions(workKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "work_key命名不规范",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.WorkConfigMap[workKey]; ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "已经存在",
		})
	}
	workName, err = reqDataJs.Get("work_name").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数work_name解析失败",
		})
	}
	workDesc, _ = reqDataJs.Get("work_desc").String()
	config, _ = reqDataJs.Get("config").String()
	if config == "" {
		config = "{}"
	}
	err = mem.CheckWorkConfig(config)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数config格式错误",
		})
	}

	// 检查flowkey是否重复
	var workConfig dal.TermiteWorkConfig
	workConfig.WorkKey = workKey
	workConfig.WorkName = workName
	workConfig.WorkDesc = workDesc
	workConfig.WorkConfig = config
	err = db.UpdateTermiteWorkConfig(workConfig)
	go mem.ForceRefresh()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "创建失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "创建成功",
	})
}

func vConfigCopyWork(ctx echo.Context) error {
	var srcWorkKey string
	var workKey string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求参数错误",
		})
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	workKey, err = reqDataJs.Get("work_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数work_key解析失败",
		})
	}
	err = utils.CheckNaminConventions(workKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "work_key命名不规范",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.WorkConfigMap[workKey]; ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "work已存在",
		})
	}
	srcWorkKey, err = reqDataJs.Get("src_work_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数src_work_key解析失败",
		})
	}
	//log.Println("workConfigMap:", mem.GlobalTermiteConfig.WorkConfigMap)
	if _, ok := mem.GlobalTermiteConfig.WorkConfigMap[srcWorkKey]; !ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "work不存在",
		})
	}
	dwork, err := db.GetTermiteWorkConfig(srcWorkKey)
	var workConfig dal.TermiteWorkConfig
	workConfig.WorkConfig = dwork.WorkConfig
	workConfig.WorkName = dwork.WorkName
	workConfig.WorkDesc = dwork.WorkDesc
	workConfig.WorkKey = workKey
	err = db.UpdateTermiteWorkConfig(workConfig)
	go mem.ForceRefresh()
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "复制失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "复制成功",
	})
}

func vDeleteWorkConfig(ctx echo.Context) error {
	var workKey string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求参数错误",
		})
	}
	defer func() {
		_ = reqBodyR.Close()
	}()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	workKey, _ = reqDataJs.Get("work_key").String()
	err = utils.CheckNaminConventions(workKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "work_key格式错误",
		})
	}
	db.DeleteTermiteWorkConfig(workKey)
	go mem.ForceRefresh()
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "删除成功",
	})
}

func vConfigUpdateWork(ctx echo.Context) error {
	var workKey string
	var workDesc string
	var workName string
	var config string
	reqBodyR := ctx.Request().Body
	if reqBodyR == nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求参数错误",
		})
	}
	defer reqBodyR.Close()
	var reqBody []byte
	var err error
	reqBody, err = ioutil.ReadAll(reqBodyR)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	reqDataJs, err := simplejson.NewJson([]byte(reqBody))
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "请求解析失败",
		})
	}
	workKey, err = reqDataJs.Get("work_key").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数错误:work_key",
		})
	}
	err = utils.CheckNaminConventions(workKey)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数不规范:work_key",
		})
	}
	if _, ok := mem.GlobalTermiteConfig.WorkConfigMap[workKey]; !ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "N0001",
			"msg":  "工作不存在",
		})
	}
	workName, err = reqDataJs.Get("work_name").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数错误:work_name",
		})
	}
	workDesc, _ = reqDataJs.Get("work_desc").String()
	config, err = reqDataJs.Get("config").String()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数错误:config",
		})
	}
	if config == "" {
		config = "{}"
	}
	err = mem.CheckWorkConfig(config)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": "P0001",
			"msg":  "参数错误:config, " + string(err.Error()),
		})
	}
	workConfig, err := db.GetTermiteWorkConfig(workKey)
	workConfig.WorkKey = workKey
	workConfig.WorkDesc = workDesc
	workConfig.WorkConfig = config
	workConfig.WorkName = workName
	err = db.UpdateTermiteWorkConfig(workConfig)
	go mem.ForceRefresh()
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"code": "P0001",
			"msg":  "更新失败",
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"code": "A0001",
		"msg":  "更新成功",
	})
}
