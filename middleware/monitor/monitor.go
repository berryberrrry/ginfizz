/*
 * @Author: berryberry
 * @LastAuthor: Do not edit
 * @since: 2019-05-10 19:28:56
 * @lastTime: 2019-06-03 21:18:40
 */
package monitor

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"route", "method"},
)

var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"route", "method"},
)

var httpQueuedRequestCount = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_queued_request_count",
		Help: "http queued request count",
	},
	[]string{"route", "method"},
)

var httpProcessingRequestCount = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_processing_request_count",
		Help: "http processing request count",
	},
	[]string{"route", "method"},
)

func init() {
	prometheus.MustRegister(
		httpRequestCount,
		httpRequestDuration,
		httpQueuedRequestCount,
		httpProcessingRequestCount,
	)
}

func Metric(c *gin.Context) {
	startTime := time.Now()
	path := c.Request.URL.String()
	method := c.Request.Method
	httpRequestCount.WithLabelValues(path, method).Inc()
	n := rand.Intn(100)
	if n >= 95 {
		time.Sleep(100 * time.Millisecond)
	} else {
		time.Sleep(50 * time.Millisecond)
	}

	elapsed := (float64)(time.Since(startTime) / time.Millisecond)
	httpRequestDuration.WithLabelValues(path, method).Observe(elapsed)

	c.Next()
}

func MaxAllowedAndMontiorQueuedProcessingRequest(n int) gin.HandlerFunc {
	if n < 0 {
		return func(c *gin.Context) {}
	}
	sem := make(chan struct{}, n)

	acquire := func(path, method string) {
		sem <- struct{}{}
		// 退出等待队列

		httpQueuedRequestCount.WithLabelValues(path, method).Dec()
		// 进入运行队列
		httpProcessingRequestCount.WithLabelValues(path, method).Inc()
	}
	release := func(path, method string) {
		<-sem
		// 请求完成，退出运行队列
		httpProcessingRequestCount.WithLabelValues(path, method).Dec()
	}
	return func(c *gin.Context) {
		path := c.Request.URL.String()
		method := c.Request.Method
		// 进入等待队列
		httpQueuedRequestCount.WithLabelValues(path, method).Inc()

		acquire(path, method)
		defer release(path, method)
		c.Next()
	}
}
