package main

import "github.com/diaohaha/termite/pkg/termite_client"

func main() {
	//err := sentry.Init(sentry.ClientOptions{
	//	Dsn: dal.Env.Sentry_Dsn,
	//})
	//
	//if err != nil {
	//	fmt.Printf("Sentry initialization failed: %v\n", err)
	//}
	//f, err := os.Open("filename.ext")
	//if err != nil {
	//	sentry.CaptureException(err)
	//	sentry.Flush(time.Second * 5)
	//} else {
	//	log.Println(f)
	//}
	//db.InitDB()
	//var configStr = "{\"works\":[\"w1\",\"w2\"],\"dags\":{\"w1\":{\"dependences\":[\"w2\"],\"trigger_rule\":\"one_done\"}}}"
	//configStr := "{\"works\":[\"w1\",\"w2\"],\"dags\":{\"w1\":{\"dependences\":[], \"trigger_rule\":\"one_done\"}}}"
	//log.Println(time.Now().Unix())
	//configStr := "{\"works\":\"w1\"}"
	//log.Println(db.CheckFlowConfig(configStr))
	//err := utils.CheckNaminConventions("1A0aab00")
	//err := db.CheckWorkConfig(configStr)
	//log.Println("check: ", err)
	//db.InitDB()
	//mem.InitMem()
	//var workIds = map[string]int64{}
	//log.Println(len(workIds))
	//type Hand struct {
	//	name string
	//}
	//type Dag struct {
	//	name  string
	//	hands []*Hand
	//}
	//var dag Dag
	//dag.name = "bob"
	//dag.hands = []*Hand{{name: "dag1's hand"}}
	//
	//// deep copy dag
	//dag2 := Dag{
	//	name: dag.name,
	//}
	//for _, hand := range dag.hands {
	//	var tmpHand Hand
	//	tmpHand = *hand
	//	dag2.hands = append(dag2.hands, &tmpHand)
	//}
	////dag2 := dag
	//
	//dag.name = "alice"
	//dag2.hands[0].name = "dag2's hand"
	//fmt.Println("dog2 name: ", dag2.name)
	//fmt.Println("dog2 hand: ", dag2.hands[0].name)
	//fmt.Println("dog1 name: ", dag.name)
	//fmt.Println("dog1 hand: ", dag.hands[0].name)

	//stable.InitSentry()
	//log.Println(dal.Env.IsDebug)
	//log.Println(raven.URL())
	//stable.CaptureError(errors.New("test sentry"), "Mem-Config", "refreshConfig", map[string]string{}, map[string]string{})
	//raven.CaptureError(errors.New("111212"), map[string]streng{})
	//time.Sleep(time.Second * 5)
	//db.InitDB()
	//mem.InitMem()
	//paritions := []int{1}
	//scheduler.WorkRecover(paritions)
	//redis.Init()
	//a := int64(1212121)
	//err := redis.Lock(a)
	//log.Println(err)
	//i, err := redis.Conn.SetNX("hello0", "12", 3*time.Second).Result()
	//log.Println(i)
	//log.Println("error:", err)
	//fmt.Println(3 / 2)
	//db.InitDB()
	//mem.InitMem()
	//scheduler.WorkTimeoutRetry()
	//scheduler.WorkTimeoutCheck()
	//db.SetWorkDelay(4998, 10)
	//a := fmt.Sprintf("%%s%", "a")
	//fmt.Println(a)
	termite_client.AddWorkFlow("124", "qa", "find_group_chat_song")
}
