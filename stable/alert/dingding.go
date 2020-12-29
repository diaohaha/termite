package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unsafe"
)

// send to dingding

const (
	//钉钉群机器人webhook地址
	dingdingURL= ""
	//dingdingURL = "https://oapi.dingtalk.com/robot/send?access_token=145d10f692bffe2e5a31050da10860ba36b64717a33ff273a42a0d34682dfe15"
	//dingdingURL = "https://oapi.dingtalk.com/robot/send?access_token=bbb5ccc0e22b11de193b83cf1f02cd724fa96cecccdd7b17cd20bd0800899a11"
)

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvator     string `json:"hideAvator"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleUrl"`
}

type ActionCardReq struct {
	ActionCard ActionCard `json:"actionCard"`
	Msgtype    string     `json:"msgtype"`
}

func SendDingMsg(alert TermiteAlertError) (err error) {
	text := markDownFormat(alert)
	ac := ActionCard{
		Title:          alert.Message,
		Text:           text,
		HideAvator:     "0",
		BtnOrientation: "0",
		SingleTitle:    "查看详情",
		SingleURL:      alert.Url,
	}
	request := ActionCardReq{
		ActionCard: ac,
		Msgtype:    "actionCard",
	}
	log.Println(request)
	b, err := json.Marshal(request)
	if err != nil {
		log.Println("jsonerror", err)
	}
	log.Println("b", string(b))

	strreq := (*string)(unsafe.Pointer(&b))
	log.Println("str req:", *strreq)
	//log.Println(requestBody)
	r := bytes.NewReader(b)
	req, err := http.NewRequest("POST", dingdingURL, r)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	log.Println(*str)
	return
}

func markDownFormat(alert TermiteAlertError) (text string) {
	text += "\n\n<br/>"
	text = text + "![监控报警]()"
	text = text + "<font  color=#A52A2A size=1>" + alert.Message + "</font>" + "\n\n<br/>"
	text = text + "<font color=gray size=72>故障时间: </font>" + "<font color=#000000 size=72>" + alert.StartTime.Format("2006-01-02 03:04:05") + " </font>" + "\n\n<br/>"
	text = text + "<font color=gray size=72>检测项: </font>" + "<font color=#000000  size=72>" + alert.Item + " </font>" + "\n\n<br/>"
	for k, v := range alert.Info {
		text = text + "<font color=gray size=72>" + k + ": </font>" + "<font color=#000000 size=72>" + v + " </font>" + "\n\n<br/>"
	}
	if len(alert.Detail) > 100 {
		text = text + "<font size=23 color=#00688B>" + strings.Replace(alert.Detail, "\"", "\\\"", 0)[:100] + "... </font>" + "\n\n<br/>"
	} else {
		text = text + "<font size=23 color=#00688B>" + strings.Replace(alert.Detail, "\"", "\\\"", 0) + " </font>" + "\n\n<br/>"
	}
	return text
}
