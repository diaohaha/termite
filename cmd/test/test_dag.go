package main

import (
	Dag "github.com/diaohaha/termite/dag/model"
	"log"
)

func main() {
	test1()
}

func test1() {
	tmp := []Dag.DagConfig{}
	tmp = append(tmp, Dag.DagConfig{Name: "1", Dependences: []string{}})
	tmp = append(tmp, Dag.DagConfig{Name: "2", Dependences: []string{"1"}})
	tmp = append(tmp, Dag.DagConfig{Name: "3", Dependences: []string{"1"}})
	tmp = append(tmp, Dag.DagConfig{Name: "4", Dependences: []string{"2", "3"}})
	tmp = append(tmp, Dag.DagConfig{Name: "5", Dependences: []string{"1"}})
	tmp = append(tmp, Dag.DagConfig{Name: "6", Dependences: []string{"4", "5"}})
	tmp = append(tmp, Dag.DagConfig{Name: "7", Dependences: []string{}})
	root := Dag.InitDag(&tmp)
	if root == nil {
		panic("Not DAG")
	}
	/*
	          1 （root)                  7(root)
	          |
	   --------------
	   |     |      |
	   ▼     ▼      |
	   2     3      |
	   |     |      ▼
	   -------      5
	       |        |
	       ▼        |
	       4        |
	       |        |
	       ----------
	           |
	           ▼
	           6
	*/
	var readys []string
	//readys = root.GetReadyNodes()
	//for _, name := range readys {
	//	log.Println("ready1:" + name)
	//}
	//log.Println("1===================")
	workStatus := map[string]int{
		"1": Dag.STATE_FINISH,
		"2": Dag.STATE_FINISH,
		"3": Dag.STATE_FINISH,
		"4": Dag.STATE_FINISH,
		"5": Dag.STATE_WAITING,
		"6": Dag.STATE_FINISH,
		"7": Dag.STATE_WAITING,
	}

	root.UpdateStatus(workStatus)

	readys = root.GetReadyNodes()
	for _, name := range readys {
		log.Println("ready2:" + name)
	}
	log.Println("2===================")

}
