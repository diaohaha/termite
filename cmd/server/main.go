package main

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/dal/prometheus"
	"github.com/diaohaha/termite/dal/redis"
	"github.com/diaohaha/termite/handler"
	proto "github.com/diaohaha/termite/proto"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	"time"
)

func init() {
	// sentry config
	//err := sentry.Init(sentry.ClientOptions{
	//	Dsn: dal.Env.Sentry_Dsn,
	//})
	//if err != nil {
	//	fmt.Printf("Sentry initialization failed: %v\n", err)
	//}

	redis.Init()
	db.InitDB()
	mem.InitMem()
}

func main() {
	// Create a new service. Optionally include some options here.

	// step2: Registry Service Handler
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			dal.Env.MicroRegistry,
		}
	})

	_ = transport.DefaultTransport.Init()
	service := micro.NewService(
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Name("termite"),
		micro.Registry(reg),
	)
	service.Init()
	s := service.Server()
	_ = s.Init(
		server.Address("0.0.0.0:17002"),
	)
	println(s.Options().Address)
	_ = proto.RegisterTermiteHandler(s, &handler.TermiteHandler{})

	go startHttp(":10223")

	// step3: Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startHttp(address string) {
	prometheus.Init()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/metrics/", echo.WrapHandler(prometheus.GetMonitorHandle()))

	e.Logger.Fatal(e.Start(address))
}
