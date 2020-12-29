package mem

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"strconv"
)

func ParseFlowConfig(flowConfig dal.TermiteFlowConfig) (flowCacheConfig CacheTermiteFlowConfig, err error) {
	// parse flow config
	// env
	var flowEnv = make(map[string]string)
	err = json.Unmarshal([]byte(flowConfig.Env), &flowEnv)
	if err != nil {
		stable.CaptureError(err, "Unmarshal", "ParseFlowConfig", map[string]string{}, map[string]string{
			"method": "ParseFlowConfig",
		})
		return
	}

	// config
	var dagConfig DagConfig
	dagJson, err := simplejson.NewJson([]byte(flowConfig.FlowConfig))
	if err != nil {
		stable.CaptureError(err, "Unmarshal", "ParseFlowConfig", map[string]string{}, map[string]string{
			"method": "ParseFlowConfig",
		})
		return
	}
	var maxSchedulerCount int
	maxSchedulerCount, err = dagJson.Get("max_scheduler_count").Int()
	if err != nil {
		err = nil
		maxSchedulerCount = DEFAULT_MAX_LIMIT
	}
	if maxSchedulerCount > MAX_SCHEDULER_LIMIT {
		// 限制最大调度量
		dagConfig.MaxSchedulerCount = MAX_SCHEDULER_LIMIT
	} else {
		dagConfig.MaxSchedulerCount = maxSchedulerCount
	}

	var schedulerMode string
	schedulerMode, err = dagJson.Get("scheduler_mode").String()
	if err != nil {
		err = nil
		schedulerMode = dal.SCHEDULER_MODE_FIFO
	} else {
		switch schedulerMode {
		case dal.SCHEDULER_MODE_FIFO:
		case dal.SCHEDULER_MODE_LIFO:
		default:
			schedulerMode = dal.SCHEDULER_MODE_FIFO
			stable.CaptureError(err, "schedulerMode", "ParseFlowConfig", map[string]string{}, map[string]string{
				"method":        "ParseFlowConfig",
				"message":       "unknow scheduler_mode",
				"schedulerMode": schedulerMode,
			})
		}
	}
	dagConfig.SchedulerMode = schedulerMode

	works, err := dagJson.Get("works").StringArray()
	if err != nil {
		stable.CaptureError(err, "", "ParseFlowConfig", map[string]string{}, map[string]string{
			"method":  "ParseFlowConfig",
			"message": "get works error!",
		})
		return
	}
	for _, work := range works {
		dagConfig.Works = append(dagConfig.Works, work)
	}
	dags, err := dagJson.Get("dags").Map()
	if err != nil {
		stable.CaptureError(err, "", "ParseFlowConfig", map[string]string{}, map[string]string{
			"method":  "ParseFlowConfig",
			"message": "get dags error!",
		})
		return
	}
	dagWorksMap := make(map[string]DagWorkConfig)
	for workName, dag := range dags {
		var dagWorkConfig DagWorkConfig
		dependences := dag.(map[string]interface{})["dependences"].([]interface{})
		for _, dependence := range dependences {
			dagWorkConfig.Dependences = append(dagWorkConfig.Dependences, dependence.(string))
		}
		dagWorkConfig.TriggerRule = dag.(map[string]interface{})["trigger_rule"].(string)
		dagWorksMap[workName] = dagWorkConfig
	}
	dagConfig.Dags = dagWorksMap

	flowCacheConfig = CacheTermiteFlowConfig{
		Key:    flowConfig.FlowKey,
		Name:   flowConfig.FlowName,
		Desc:   flowConfig.FlowDesc,
		Env:    flowEnv,
		Config: dagConfig,
		Switch: flowConfig.Switch,
	}
	return
}

func ParseWorkConfig(workConfig dal.TermiteWorkConfig) (workCacheConfig CacheTermiteWorkConfig, err error) {
	// parse work config
	var workitsConfig = make(map[string]string)
	var execTimeout int64
	var pushTimeout int64
	var retries int
	var delaies int
	err = json.Unmarshal([]byte(workConfig.WorkConfig), &workitsConfig)
	if err != nil {
		stable.CaptureError(err, "", "ParseWorkConfig", map[string]string{}, map[string]string{
			"method":  "ParseWorkConfig",
			"message": "Unmarshal error!",
		})
	}
	if _, ok := workitsConfig["exec_timeout"]; ok {
		execTimeout, err = strconv.ParseInt(workitsConfig["exec_timeout"], 10, 64)
		if err != nil {
			execTimeout = DEFAULT_EXEC_TIMEOUT
		}
	} else {
		execTimeout = DEFAULT_EXEC_TIMEOUT
	}
	if _, ok := workitsConfig["push_timeout"]; ok {
		pushTimeout, err = strconv.ParseInt(workitsConfig["push_timeout"], 10, 64)
		if err != nil {
			pushTimeout = DEFAULT_PUSH_TIMEOUT
		}
	} else {
		pushTimeout = DEFAULT_PUSH_TIMEOUT
	}
	if _, ok := workitsConfig["retries"]; ok {
		retries, err = strconv.Atoi(workitsConfig["retries"])
		if err != nil {
			retries = DEFAULT_RETRIES
		}
		if retries > MAX_RETRIES {
			retries = MAX_RETRIES
		}
	} else {
		retries = DEFAULT_RETRIES
	}
	if _, ok := workitsConfig["delaies"]; ok {
		delaies, err = strconv.Atoi(workitsConfig["delaies"])
		if err != nil {
			delaies = DEFAULT_DELAIES
		}
		if delaies > MAX_DELAIES {
			delaies = MAX_DELAIES
		}
	} else {
		delaies = DEFAULT_RETRIES
	}
	workCacheConfig = CacheTermiteWorkConfig{
		Key:         workConfig.WorkKey,
		Name:        workConfig.WorkName,
		Desc:        workConfig.WorkDesc,
		Config:      workitsConfig,
		PushTimeout: pushTimeout,
		ExecTimeout: execTimeout,
		Retries:     retries,
		Delaies:     delaies,
	}
	return
}
