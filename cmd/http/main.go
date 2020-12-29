package main

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/router"
)

func main() {
	// Create a new service. Optionally include some options here.
	// sentry config
	//go func() { router.StartHttp("0:0:0:0:10023") }()
	db.InitDB()
	mem.InitMem()
	router.StartHttp(":10023")
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, os.Kill)
}
