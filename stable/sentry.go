package stable

import (
	"github.com/diaohaha/termite/dal"
	"github.com/getsentry/raven-go"
	"log"
	"time"
)

//var SentryClient *raven.Client

func InitSentry() {
	var err error
	var SentryDSN string
	if dal.Env.IsDebug {
		SentryDSN = "***"
	} else {
		SentryDSN = "***"
	}
	//SentryClient, err = raven.New(SentryDSN)
	err = raven.SetDSN(SentryDSN)
	if err != nil {
		panic(err)
	}
}

func CaptureErrorWithTime(err error, endpoint, method string, req interface{}, sentryCapture map[string]string, t time.Time) {
	raven.CaptureError(err, sentryCapture)
}

func CaptureError(err error, endpoint, method string, req interface{}, sentryCapture map[string]string) {
	log.Println("Exception: ", err, "endpoint: ", endpoint, "method: ", method)
	eventId := raven.CaptureError(err, sentryCapture)
	log.Println("eventId:", eventId)
}
