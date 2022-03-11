package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/leonardonatali/metrics/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := prometheus.NewRegistry()

	r.MustRegister(metrics.OnlineUsers)
	r.MustRegister(metrics.HTTPRequestsTotal)
	r.MustRegister(metrics.HTTPDuration)

	d := promhttp.InstrumentHandlerDuration(
		metrics.HTTPDuration.MustCurryWith(prometheus.Labels{"path": "/"}),
		promhttp.InstrumentHandlerCounter(metrics.HTTPRequestsTotal, http.HandlerFunc(home)),
	)

	http.Handle("/", promhttp.InstrumentHandlerCounter(metrics.HTTPRequestsTotal, d))

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			metrics.OnlineUsers.Set(float64(rand.Intn(2000)))
		}
	}()

	log.Fatal(http.ListenAndServe(":8181", nil))
}

func home(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(4500)) * time.Millisecond)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf("Hello, World!\nIs %s\n", time.Now().Format("2006-01-02 15:04:05"))))
}
