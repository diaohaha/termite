package main

import (
	"github.com/diaohaha/termite/dal/db"
	"github.com/diaohaha/termite/dal/mem"
	"github.com/diaohaha/termite/stable/alert"
)

//type TermiteAlertError struct {
//	Message   string            `json:"message"`
//	Item      string            `json:"item"`
//	Result    string            `json:"result"`
//	Level     int               `json:"level"`
//	StartTime time.Time         `json:"start_time"`
//	Info      map[string]string `json:"info"`
//	Url       string            `json:"url"`
//}

func main() {
	//a := alert.TermiteAlertError{
	//	Message:   "flow: video_release receive 676 work error in 5s",
	//	Item:      "短时间大量异常",
	//	Url:       "http://www.baidu.com",
	//	StartTime: time.Now(),
	//	Info:      map[string]string{},
	//}
	db.InitDB()
	mem.InitMem()
	if a, ok := alert.CheckReceiveLargeWorkFail("low_h264_encode_submit_task"); ok {
		//log.Println(a)
		alert.SendDingMsg(a)
	}
}
