package main

import (
	"github.com/diaohaha/termite/dag"
	"github.com/diaohaha/termite/dag/model"
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/scheduler"
	"log"
)

func main() {
	db.InitDB()
	mem.InitMem()
	scheduler.Init()
	log.Println(*scheduler.GlobalflowDagGraphMap["video_dynamic_cover"])
	mdag := scheduler.GlobalflowDagGraphMap["video_dynamic_cover"]
	copydag, _ := dag.DeepCopyDagNode(*mdag, map[string]*model.DagNode{})
	log.Println(mdag)
	log.Println(mdag.Children)
	log.Println(copydag)
	log.Println(copydag.Children)
}
