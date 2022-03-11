package metrics

import "github.com/prometheus/client_golang/prometheus"

var HTTPRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "go_app_http_requests_total",
	Help: "Count of all HTTP requests",
}, []string{})

var HTTPDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "go_app_http_duration",
	Help: "Duration in seconds",
}, []string{"path"})
