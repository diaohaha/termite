package mem

import (
	"bytes"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/stable"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func refreshConfig() {
	// TempGlobalTermiteConfig
	log.Println("refresh mem configs ...")
	workConfigs, err := db.GetTermiteWorkConfigs()
	if err != nil {
		stable.CaptureError(err, "Mem-Config", "refreshConfig", map[string]string{}, map[string]string{})
	}
	flowConfigs, err := db.GetTermiteFlowConfigs()
	if err != nil {
		stable.CaptureError(err, "Mem-Config", "refreshConfig", map[string]string{}, map[string]string{})
	}

	TempGlobalTermiteConfig.WorkConfigMap = map[string]CacheTermiteWorkConfig{}
	TempGlobalTermiteConfig.FlowConfigMap = map[string]CacheTermiteFlowConfig{}

	for _, workConfig := range workConfigs {
		memWorkConfig, err := ParseWorkConfig(workConfig)
		if err != nil {
			stable.CaptureError(err.(error), "refreshConfig", "refreshConfig", map[string]string{}, map[string]string{
				"msg":     "parse config error!",
				"type":    "panic",
				"work":    workConfig.WorkName,
				"workKey": workConfig.WorkKey,
			})
		} else {
			TempGlobalTermiteConfig.WorkConfigMap[workConfig.WorkKey] = memWorkConfig
		}
	}

	for _, flowConfig := range flowConfigs {
		memFlowConfig, err := ParseFlowConfig(flowConfig)
		if err != nil {
			stable.CaptureError(err.(error), "refreshConfig", "refreshConfig", map[string]string{}, map[string]string{
				"msg":     "parse config error!",
				"type":    "panic",
				"flow":    flowConfig.FlowName,
				"flowKey": flowConfig.FlowKey,
			})
		} else {
			TempGlobalTermiteConfig.FlowConfigMap[flowConfig.FlowKey] = memFlowConfig
		}
	}
	TempGlobalTermiteConfig.ExpireTimeStamp = time.Now().Unix() + EXPIRETIME
	GlobalTermiteConfig = TempGlobalTermiteConfig
}

func printConfig() {
	fmt.Println("********************** 工作 ***********************")
	for workKey, workConfig := range GlobalTermiteConfig.WorkConfigMap {
		fmt.Println(workKey, ":")
		confBytes, _ := json.Marshal(workConfig)
		var out bytes.Buffer
		_ = json.Indent(&out, confBytes, "", "\t")
		_, _ = out.WriteTo(os.Stdout)
	}
	fmt.Println("********************** 工作流 ***********************")
	for flowKey, flowConfig := range GlobalTermiteConfig.FlowConfigMap {
		fmt.Println(flowKey, ":")
		confBytes, _ := json.Marshal(flowConfig)
		var out bytes.Buffer
		_ = json.Indent(&out, confBytes, "", "\t")
		_, _ = out.WriteTo(os.Stdout)
	}
	fmt.Println("****************************************************")
}

func ForceRefresh() {
	log.Println("force refresh ...")
	refreshConfig()
}

func TryRefresh() {
	log.Println("try refresh ...")
	if time.Now().Unix() > GlobalTermiteConfig.ExpireTimeStamp {
		refreshConfig()
	}
}

// Query

func GetFlowConfigMap() map[string]CacheTermiteFlowConfig {
	if time.Now().Unix() > GlobalTermiteConfig.ExpireTimeStamp {
		refreshConfig()
	}
	return GlobalTermiteConfig.FlowConfigMap
}

func GetFlowConfig(flowKey string) CacheTermiteFlowConfig {
	if time.Now().Unix() > GlobalTermiteConfig.ExpireTimeStamp {
		refreshConfig()
	}
	return GlobalTermiteConfig.FlowConfigMap[flowKey]
}

func GetWorkConfigMap() map[string]CacheTermiteWorkConfig {
	log.Println(GlobalTermiteConfig.ExpireTimeStamp)
	if time.Now().Unix() > GlobalTermiteConfig.ExpireTimeStamp {
		refreshConfig()
	}
	return GlobalTermiteConfig.WorkConfigMap
}

func GetWorkConfig(workKey string) CacheTermiteWorkConfig {
	if time.Now().Unix() > GlobalTermiteConfig.ExpireTimeStamp {
		refreshConfig()
	}
	return GlobalTermiteConfig.WorkConfigMap[workKey]
}
