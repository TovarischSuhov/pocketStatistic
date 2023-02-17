package main

import (
	"net/http"

	"github.com/TovarischSuhov/go-template/internal/handlers"
	"github.com/TovarischSuhov/go-template/internal/metrics"
	"github.com/TovarischSuhov/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.LogLevel = log.DebugLevel
	log.UseColors = true
	go metrics.CountMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/random", handlers.GetRandomTopicHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error(err.Error())
	}
}
