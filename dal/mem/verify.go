package mem

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/stable"
	"github.com/diaohaha/termite/utils"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/pingcap/errors"
	"go.uber.org/zap"
	"log"
)

func CheckFlowConfig(config string) (bool, error) {
	/*
		check flow config format
	*/
	var err error
	var hasError bool
	utils.Try(func() {
		var dagConfig DagConfig
		dagJson, err := simplejson.NewJson([]byte(config))
		if err != nil {
			stable.Logger.Error(
				"Parse Config Bytes Error!",
				zap.String("func", "CheckFlowConfig"),
				zap.Error(err),
			)
			panic(err)
		}
		stable.Logger.Info(
			"Parse Config Bytes Success!",
			zap.String("func", "CheckFlowConfig"),
		)
		works, err := dagJson.Get("works").StringArray()
		if err != nil {
			stable.Logger.Error(
				"Parse Config Bytes Error: do not have key:works",
				zap.String("func", "CheckFlowConfig"),
				zap.Error(err),
			)
			panic(err)
		}
		for _, work := range works {
			dagConfig.Works = append(dagConfig.Works, work)
		}
		dags, err := dagJson.Get("dags").Map()
		if err != nil {
			stable.Logger.Error(
				"Parse Config Bytes Error: do not have key:dags",
				zap.String("func", "CheckFlowConfig"),
				zap.Error(err),
			)
			panic(err)
		}
		dagWorksMap := make(map[string]DagWorkConfig)
		for workName, dag := range dags {
			fmt.Println("parse dag:", dag)
			var dagWorkConfig DagWorkConfig
			dependences := dag.(map[string]interface{})["dependences"].([]interface{})
			for _, dependence := range dependences {
				var dependenceWork = dependence.(string)
				var nullTwConfig dal.TermiteWorkConfig
				twConfig, err := db.GetTermiteWorkConfig(dependenceWork)
				if err != nil {
					stable.Logger.Error(
						"GetTermiteWorkConfig Error!",
						zap.Error(err),
					)
				}
				log.Println("twConfig", twConfig)
				if twConfig == nullTwConfig {
					panic(errors.New("work not exist"))
				}
				dagWorkConfig.Dependences = append(dagWorkConfig.Dependences, dependence.(string))
			}
			dagWorkConfig.TriggerRule = dag.(map[string]interface{})["trigger_rule"].(string)
			dagWorksMap[workName] = dagWorkConfig
		}
		dagConfig.Dags = dagWorksMap
		log.Println("thisisdag", dagConfig)
	}).CacheAll(func(e interface{}) {
		log.Println("has error:", e)
		err = e.(error)
		hasError = true
	}).Finally(func() {
	})
	return !hasError, err
}

func CheckWorkConfig(config string) error {
	type WorkConfig map[string]string
	var workConfig = make(WorkConfig)
	err := json.Unmarshal([]byte(config), &workConfig)
	return err
}
