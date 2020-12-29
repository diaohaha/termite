package router

import (
	xm "github.com/diaohaha/termite/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartHttp(address string) {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(xm.BodyParser()) // 这个中间件会把参数设置到ctx里

	// 是否启用
	//e.Use(eryptMiddleware)

	// config
	e.POST("/api/config/flow/query/", vQueryFlowConfig)
	e.POST("/api/config/flow/query/v2/", vQueryFlowConfigV2)
	e.POST("/api/config/flow/copy/", vConfigCopyFlow)
	e.POST("/api/config/flow/create/", vConfigCreateFlow)
	e.POST("/api/config/flow/update/", vConfigUpdateFlow)
	e.POST("/api/config/flow/delete/", vDeleteFlowConfig)
	e.POST("/api/config/flow/switch/", vSwitchFlowConfig)

	e.POST("/api/config/work/query/", vQueryWorkConfig)
	e.POST("/api/config/work/query/v2/", vQueryWorkConfigV2)
	e.POST("/api/config/work/copy/", vConfigCopyWork)
	e.POST("/api/config/work/create/", vConfigCreateWork)
	e.POST("/api/config/work/update/", vConfigUpdateWork)
	e.POST("/api/config/work/delete/", vDeleteWorkConfig)

	e.POST("/api/instance/flow/query/", vQueryFlowInstance)
	e.POST("/api/instance/work/query/", vQueryWorkInstance)

	e.POST("/api/instance/flow/add/", vAddFlowInstance)
	e.POST("/api/instance/flow/add/file/", vAddFlowInstanceByFile)
	e.POST("/api/instance/flow/delete/", vDeleteFlowInstance)
	e.POST("/api/instance/flow/recover/", vRecoverFlowInstance)
	e.POST("/api/instance/work/recover/", vRecoverWorkInstance)

	e.POST("/api/info/work/count/", vInfoWorkCount)
	e.POST("/api/info/flow/count/", vInfoFlowCount)

	e.File("/admin", "dist/index.html")
	e.Static("/static", "dist/static")

	// runtime

	e.Logger.Fatal(e.Start(address))
}
