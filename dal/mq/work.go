package mq

import (
	"github.com/diaohaha/termite/dal"
	"github.com/diaohaha/termite/stable"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"time"
)

/*
func SendMessage(config *rocketmq.ProducerConfig, msg string, tag string) {
	producer, err := rocketmq.NewProducer(config)

	if err != nil {
		sentry.CaptureException(err)
		fmt.Println("create Producer failed, error:", err)
		return
	}

	err = producer.Start()
	if err != nil {
		sentry.CaptureException(err)
		fmt.Println("start producer error", err)
		return
	}
	defer producer.Shutdown()

	fmt.Printf("Producer: %s started... \n", producer)
	result, err := producer.SendMessageSync(&rocketmq.Message{Topic: *topic, Body: msg, Tags: tag})
	if err != nil {
		sentry.CaptureException(err)
		fmt.Println("Error:", err)
	}
	fmt.Printf("send message: %s result: %s\n", msg, result)
}
*/

type MsgBody struct {
	Flow       string
	Work       string
	Project    string
	Work_id    int64
	Flow_id    int64
	Event      string
	Cid        string
	WorkConfig map[string]string `json:"work_config"`
	TimeStamp  int64
}

func SendTaskReadyMsg(flow string, cid string, work string, project string, work_id int64, flow_id int64, workConfig map[string]string) {
	// send termite event msg
	content := MsgBody{
		Flow:       flow,
		Work:       work,
		Work_id:    work_id,
		Flow_id:    flow_id,
		Event:      "ready",
		Project:    project,
		Cid:        cid,
		WorkConfig: workConfig,
		TimeStamp:  time.Now().UnixNano(),
	}
	log.Println("kafka host:", dal.Env.Kafka_Termite_event.Host, "topic:", dal.Env.Kafka_Termite_event.Topic)
	paramJson, err := json.Marshal(content)
	if err != nil {
		stable.CaptureError(err, "SendTaskReadyMsg", "SendTaskReadyMsg", map[string]string{}, map[string]string{
			"method":  "SendTaskReadyMsg",
			"msg":     "Marshal() Error!",
			"cid":     cid,
			"work_id": strconv.FormatInt(work_id, 10),
			"flow_id": strconv.FormatInt(flow_id, 10),
			"vflow":   flow,
			"vwork":   work,
		})
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{}
	msg.Topic = dal.Env.Kafka_Termite_event.Topic
	msg.Value = sarama.StringEncoder(paramJson)

	client, err := sarama.NewSyncProducer([]string{dal.Env.Kafka_Termite_event.Host}, config)
	if err != nil {
		stable.CaptureError(err, "SendTaskReadyMsg", "SendTaskReadyMsg", map[string]string{}, map[string]string{
			"method":  "SendTaskReadyMsg",
			"msg":     "NewSyncProducer() Error!",
			"cid":     cid,
			"work_id": strconv.FormatInt(work_id, 10),
			"flow_id": strconv.FormatInt(flow_id, 10),
			"vflow":   flow,
			"vwork":   work,
			"project": project,
		})
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		stable.CaptureError(err, "SendTaskReadyMsg", "SendTaskReadyMsg", map[string]string{}, map[string]string{
			"method":  "SendTaskReadyMsg",
			"msg":     "SendMessage() Error!",
			"cid":     cid,
			"work_id": strconv.FormatInt(work_id, 10),
			"flow_id": strconv.FormatInt(flow_id, 10),
			"vflow":   flow,
			"vwork":   work,
		})
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
