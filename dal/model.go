package dal

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	TermiteWorkStateInit    = 0
	TermiteWorkStatePush    = 1
	TermiteWorkStateRunning = 2
	TermiteWorkStateSuccess = 3
	TermiteWorkStateFailed  = 4
	TermiteWorkStateError   = 5
	TermiteWorkStateTimeout = 6
	TermiteWorkStateDelay   = 7
	TermiteWorkStateStop    = 8
)
const (
	TermiteWorkFlowStateInit    = 0
	TermiteWorkFlowStateRunning = 1
	TermiteWorkFlowStateFinish  = 2
	TermiteWorkFlowStateTimeout = 3
	TermiteWorkFlowStateError   = 4
	TermiteWorkFlowStateFailed  = 5
	TermiteWorkFlowStateDelay   = 6
)

var TermiteWorkStateTransMap = map[int32]map[string]string{
	TermiteWorkStateInit: map[string]string{
		"en": "init",
		"zh": "初始化",
	},
	TermiteWorkStatePush: map[string]string{
		"en": "push",
		"zh": "已下发",
	},
	TermiteWorkStateRunning: map[string]string{
		"en": "running",
		"zh": "执行中",
	},
	TermiteWorkStateSuccess: map[string]string{
		"en": "success",
		"zh": "成功",
	},
	TermiteWorkStateFailed: map[string]string{
		"en": "failed",
		"zh": "失败",
	},
	TermiteWorkStateError: map[string]string{
		"en": "error",
		"zh": "异常",
	},
	TermiteWorkStateTimeout: map[string]string{
		"en": "timeout",
		"zh": "超时",
	},
	TermiteWorkStateDelay: map[string]string{
		"en": "delay",
		"zh": "延期调度",
	},
	TermiteWorkStateStop: map[string]string{
		"en": "stop",
		"zh": "停止",
	},
}

var TermiteWorkFlowStateTransMap = map[int32]map[string]string{
	TermiteWorkFlowStateInit: map[string]string{
		"en": "init",
		"zh": "待调度",
	},
	TermiteWorkFlowStateRunning: map[string]string{
		"en": "running",
		"zh": "调度中",
	},
	TermiteWorkFlowStateFinish: map[string]string{
		"en": "finish",
		"zh": "完成",
	},
	TermiteWorkFlowStateTimeout: map[string]string{
		"en": "timeout",
		"zh": "超时",
	},
	TermiteWorkFlowStateError: map[string]string{
		"en": "error",
		"zh": "异常",
	},
	TermiteWorkFlowStateFailed: map[string]string{
		"en": "failed",
		"zh": "失败",
	},
	TermiteWorkFlowStateDelay: map[string]string{
		"en": "delay",
		"zh": "调度延期",
	},
}

type TermiteWork struct {
	Id        int64     `gorm:"primary_key;column:id"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	Vflow     string    `gorm:"column:vflow"`
	Vwork     string    `gorm:"column:vwork"`
	Project   string    `gorm:"column:project"`
	Vstate    int32     `gorm:"column:vstate"`
	Cid       string    `gorm:"column:cid"`
	Input     string    `gorm:"column:input"`
	Output    string    `gorm:"column:output"`
	Error     string    `gorm:"column:error"`
	Env       string    `gorm:"column:env"`
	Wakeup    int64     `gorm:"column:wake_up"`
	Retries   int       `gorm:"column:retries"`
	Delaies   int       `gorm:"column:delaies"`
}

func (TermiteWork) TableName() string {
	return "t_termite_work"
}

type TermiteFlow struct {
	Id        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	Partition int       `gorm:"column:tpartition"`
	Project   string    `gorm:"column:project"`
	Vflow     string    `gorm:"column:vflow"`
	Cid       string    `gorm:"column:cid"`
	Vstate    int32     `gorm:"column:vstate"`
	Context   string    `gorm:"column:tcontext"`
}

func (TermiteFlow) TableName() string {
	return "t_termite_workflow"
}

type TermiteFlowConfig struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdatedAt  time.Time `gorm:"column:update_time" json:"update_time"`
	FlowKey    string    `gorm:"column:flow_key" json:"flow_key"`
	FlowName   string    `gorm:"column:flow_name" json:"flow_name"`
	FlowDesc   string    `gorm:"column:flow_desc" json:"flow_desc"`
	FlowConfig string    `gorm:"column:flow_config" json:"flow_config"`
	Env        string    `gorm:"column:env" json:"env"`
	Switch     int       `gorm:"column:switch" json:"switch"`
}

func (TermiteFlowConfig) TableName() string {
	return "termite_flow_config"
}

type TermiteWorkConfig struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdatedAt  time.Time `gorm:"column:update_time" json:"update_time"`
	WorkKey    string    `gorm:"column:work_key" json:"work_key"`
	WorkName   string    `gorm:"column:work_name" json:"work_name"`
	WorkDesc   string    `gorm:"column:work_desc" json:"work_desc"`
	WorkConfig string    `gorm:"column:work_config" json:" work_config"`
}

func (TermiteWorkConfig) TableName() string {
	return "termite_work_config"
}

const (
	TRIGGER_RULE_ALL_DONE    = "all_done"
	TRIGGER_RULE_ALL_SUCCESS = "all_success"
	TRIGGER_RULE_ALL_FAIL    = "all_fail"
	TRIGGER_RULE_ONE_DONE    = "one_done"
	TRIGGER_RULE_ONE_SUCCESS = "one_success"
	TRIGGER_RULE_ONE_FAIL    = "one_fail"
)

const (
	SCHEDULER_MODE_FIFO = "fifo"
	SCHEDULER_MODE_LIFO = "lifo"
)

func GetWorkStateZh(state int32) (zh string) {
	return TermiteWorkStateTransMap[state]["zh"]
}
func GetWorkFlowStateZh(state int32) (zh string) {
	return TermiteWorkFlowStateTransMap[state]["zh"]
}

type TermiteNode struct {
	Id         int64     `gorm:"primary_key"`
	CreatedAt  time.Time `gorm:"column:create_time"`
	UpdatedAt  time.Time `gorm:"column:update_time"`
	Partitions string    `gorm:"column:partitions"`
	NodeId     string    `gorm:"column:node_id"`
	NodeType   string    `gorm:"column:node_type"`
	ExpireTime time.Time `gorm:"column:expire_time"`
}

func (TermiteNode) TableName() string {
	return "t_termite_node"
}

const (
	NODETYPE_DAG   = "dag"
	NODETYPE_DELAY = "delay"
)
