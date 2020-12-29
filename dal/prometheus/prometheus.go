package prometheus

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

var (
	MetricCounter *prometheus.CounterVec
	MetricTimer   *prometheus.HistogramVec
	//MetricVal     *prometheus.GaugeVec
)

func Init() {
	// 注册metric
	MetricCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "base_service",
			Subsystem: "termite",
			Name:      "counter",
			Help:      "count request",
		},
		[]string{"instance", "method", "endpoint", "code"},
	)
	MetricTimer = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "base_service",
			Subsystem: "termite",
			Name:      "latency",
			Help:      "consume time",
			Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5},
		},
		[]string{"instance", "method", "endpoint"},
	)

	prometheus.MustRegister(MetricCounter)
	prometheus.MustRegister(MetricTimer)
}

func GetMonitorHandle() http.Handler {
	return promhttp.Handler()
}

func HostName() string {
	name, _ := os.Hostname()
	return name
}

func Timer(method, endpoint string) *prometheus.Timer {
	t := prometheus.NewTimer(MetricTimer.WithLabelValues(HostName(), method, endpoint))
	return t
}

func counterInr(method, endpoint, code string) {
	MetricCounter.WithLabelValues(HostName(), method, endpoint, code).Inc()
}

func RouterCounterInr(endpoint, code string) {
	counterInr(TracerMethod_Router, endpoint, code)
}
func RpcCounterInr(endpoint, code string) {
	counterInr(TracerMethod_Rpc, endpoint, code)
}

func RoomUserCounterInr(roomId, showId string, count int64) {
	MetricCounter.WithLabelValues(HostName(), "room_user", roomId, showId).Add(float64(count))
}

func TimerMiddleware(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//延时
		t := Timer(TracerMethod_Router, c.Path())
		defer t.ObserveDuration()
		return handlerFunc(c)
	}
}

type tracer struct {
	t *prometheus.Timer
}

func (t *tracer) Trace() {
	t.t.ObserveDuration()
}
func IOTimeTrace(method, endpoint string) *tracer {
	return &tracer{t: Timer(method, endpoint)}
}

const (
	TracerMethod_Router     = "router"
	TracerMethod_Redis      = "redis"
	TracerMethod_Mysql      = "mysql"
	TracerMethod_Rpc        = "rpc"
	TracerMethod_Http       = "http"
)
