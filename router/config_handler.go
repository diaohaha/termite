package router

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/middleware"
	"github.com/labstack/echo"
	"log"
	"strconv"
)

func vQueryFlowConfigV2(ctx echo.Context) (err error) {
	vflow := middleware.GetStringFromCtxDefault(ctx, "vflow", "")
	search := middleware.GetStringFromCtxDefault(ctx, "search", "")
	pageSize := middleware.GetInt64FromCtxDefault(ctx, "page_size", 10)
	pageIndex := middleware.GetInt64FromCtxDefault(ctx, "page_index", 0)

	configs, count, err := db.QueryTermiteFlowConfigs(vflow, search, int(pageIndex), int(pageSize))
	retConfigs := make([]mem.CacheTermiteFlowConfig, 0)
	for _, config := range configs {
		rconfig, ierr := mem.ParseFlowConfig(config)
		if ierr != nil {
			// TODO
		} else {
			retConfigs = append(retConfigs, rconfig)
		}
	}
	return middleware.ApiResult(ctx, err, "A0001", "", map[string]interface{}{
		"count":   count,
		"configs": retConfigs,
	})
}

func vSwitchFlowConfig(ctx echo.Context) (err error) {
	vflow := middleware.GetStringFromCtxDefault(ctx, "vflow", "")
	iswitch := middleware.GetStringFromCtxDefault(ctx, "switch", "")
	intSwitch, _ := strconv.Atoi(iswitch)
	err = db.UpdateFlowSwitch(vflow, intSwitch)
	if err != nil {
		log.Println("hello", err)
		return err
	} else {
		return middleware.ApiResult(ctx, err, "A0001", "", map[string]interface{}{})
	}
}

func vQueryWorkConfigV2(ctx echo.Context) (err error) {
	vwork := middleware.GetStringFromCtxDefault(ctx, "vwork", "")
	search := middleware.GetStringFromCtxDefault(ctx, "search", "")
	pageSize := middleware.GetInt64FromCtxDefault(ctx, "page_size", 10)
	pageIndex := middleware.GetInt64FromCtxDefault(ctx, "page_index", 0)

	configs, count, err := db.QueryTermiteWorkConfigs(vwork, search, int(pageIndex), int(pageSize))
	retConfigs := make([]mem.CacheTermiteWorkConfig, 0)
	for _, config := range configs {
		rconfig, ierr := mem.ParseWorkConfig(config)
		if ierr != nil {
			// TODO
		} else {
			retConfigs = append(retConfigs, rconfig)
		}
	}
	return middleware.ApiResult(ctx, err, "A0001", "", map[string]interface{}{
		"count":   count,
		"configs": retConfigs,
	})
}
