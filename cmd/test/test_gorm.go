package main

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/dal/db"
	"fmt"
	"log"
)

func main() {
	// step1: DB connection
	fmt.Println("hello. test.")
	//var tc db.TermiteConfig
	//flows := tc.GetFlows(wdb)
	//for _, flow := range flows {
	//	println(flow)
	//}
	db.InitDB()
	//var twork dal.TermiteWork
	//err := db.WTDB.Where("id = ?", 1).First(&twork).Error
	//if err == gorm.ErrRecordNotFound {
	//	log.Println("not exist")
	//}
	//db.WTDB.Model(&twork).Update("vstate", dal.TermiteWorkStatePush)
	flowId := 630
	state := 1
	err := db.WTDB.Debug().Model(dal.TermiteFlow{}).Where("id = ?", flowId).Updates(dal.TermiteFlow{Vstate: int32(state), Context: "123"}).Error
	log.Println(err)

}
