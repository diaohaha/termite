package mem

import (
	"sync"
	"time"
)

/*
   termite config info, load to cache, refresh every 30s
*/

const (
	EXPIRETIME           = 30
	DEFAULT_MAX_LIMIT    = 0
	MAX_SCHEDULER_LIMIT  = 2000
	DEFAULT_RETRIES      = 1
	MAX_RETRIES          = 3
	DEFAULT_DELAIES      = 10
	MAX_DELAIES          = 50
	DEFAULT_EXEC_TIMEOUT = 60 * 5
	DEFAULT_PUSH_TIMEOUT = 60 * 60 * 24 // 一天
)

type DagWorkConfig struct {
	Dependences []string `json:"dependences"`
	TriggerRule string   `json:"trigger_rule"`
}

type DagConfig struct {
	Works             []string                 `json:"works"`
	Dags              map[string]DagWorkConfig `json:"dags"`
	MaxSchedulerCount int                      `json:"max_scheduler_count"`
	SchedulerMode     string                   `json:"scheduler_mode"`
}

type CacheTermiteWorkConfig struct {
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	Desc        string            `json:"desc"`
	Config      map[string]string `json:"config"`
	ExecTimeout int64             `json:"exec_timeout"`
	PushTimeout int64             `json:"push_timeout"`
	Retries     int               `json:"retries"`
	Delaies     int               `json:"delaies"`
}

type CacheTermiteFlowConfig struct {
	Key    string            `json:"key"`
	Name   string            `json:"name"`
	Desc   string            `json:"desc"`
	Env    map[string]string `json:"env"`
	Config DagConfig         `json:"config"`
	Switch int               `json:"switch"`
}

type CacheTermiteConfig struct {
	ExpireTimeStamp int64
	FlowConfigMap   map[string]CacheTermiteFlowConfig
	WorkConfigMap   map[string]CacheTermiteWorkConfig
	ConfigLock      sync.RWMutex
}

var GlobalTermiteConfig CacheTermiteConfig
var TempGlobalTermiteConfig CacheTermiteConfig

/*  mem data structure */

func InitMem() {
	refreshConfig()
	//printConfig()
	go func() {
		for true {
			time.Sleep(time.Second * 10)
			TryRefresh()
		}
	}()

}
