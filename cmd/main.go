package main

import (
	"net/http"
	"time"

	"github.com/TovarischSuhov/go-template/internal/client"
	"github.com/TovarischSuhov/go-template/internal/config"
	"github.com/TovarischSuhov/go-template/internal/domain"
	"github.com/TovarischSuhov/log"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.LogLevel = log.DebugLevel
	log.UseColors = true
	conf := config.GetConfig()
	unreadGauge := prom.NewGauge(prom.GaugeOpts{Name: "unread"})
	unreadTimeGauge := prom.NewGauge(prom.GaugeOpts{Name: "unread_time"})
	archivedGauge := prom.NewGauge(prom.GaugeOpts{Name: "archived"})
	archivedTimeGauge := prom.NewGauge(prom.GaugeOpts{Name: "archived_time"})
	prom.MustRegister(unreadGauge)
	prom.MustRegister(unreadTimeGauge)
	prom.MustRegister(archivedGauge)
	prom.MustRegister(archivedTimeGauge)
	go func() {
		c := client.NewHTTPClient(conf.ConsumerKey, conf.AccessToken, conf.Host, conf.Path)
		for {
			result, err := c.GetTopicsList()
			if err != nil {
				log.Error(err.Error())
			}
			var (
				unreadCount   int64
				archivedCount int64
				unreadTime    int64
				archivedTime  int64
			)
			for _, topic := range result.List {
				switch topic.Status {
				case domain.UNREAD_STATUS:
					unreadCount++
					unreadTime += topic.TimeToRead
				case domain.ARCHIVED_STATUS:
					archivedCount++
					archivedTime += topic.TimeToRead
				}
			}
			log.Info("unread: %d, unread_time: %d, archived: %d, archived_time: %d", unreadCount, unreadTime, archivedCount, archivedTime)
			unreadGauge.Set(float64(unreadCount))
			unreadTimeGauge.Set(float64(unreadTime))
			archivedGauge.Set(float64(archivedCount))
			archivedTimeGauge.Set(float64(archivedTime))
			time.Sleep(time.Hour)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error(err.Error())
	}
}
