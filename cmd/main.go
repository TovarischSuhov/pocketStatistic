package main

import (
	"net/http"
	"time"

	"github.com/TovarischSuhov/go-template/internal/handlers"
	"github.com/TovarischSuhov/go-template/internal/metrics"
	"github.com/TovarischSuhov/log"
	"github.com/go-co-op/gocron/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.LogLevel = log.DebugLevel
	log.UseColors = true
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Error(err.Error())
	}
	defer func() { s.Shutdown() }()
	_, err = s.NewJob(
		gocron.DurationJob(time.Minute*15),
		gocron.NewTask(metrics.CountMetrics),
	)
	if err != nil {
		log.Error(err.Error())
	}
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/random", handlers.GetRandomTopicHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error(err.Error())
	}
}
