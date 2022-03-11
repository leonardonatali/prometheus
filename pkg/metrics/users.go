package metrics

import "github.com/prometheus/client_golang/prometheus"

var OnlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "go_app_online_users",
	Help: "Online users",
	ConstLabels: prometheus.Labels{
		"page": "home",
	},
})
