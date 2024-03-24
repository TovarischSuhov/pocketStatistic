package metrics

import (
	"math/rand"
	"sync"
	"time"

	"github.com/TovarischSuhov/go-template/internal/client"
	"github.com/TovarischSuhov/go-template/internal/config"
	"github.com/TovarischSuhov/go-template/internal/domain"
	"github.com/TovarischSuhov/log"
	prom "github.com/prometheus/client_golang/prometheus"
)

var (
	topicsGauge   *prom.GaugeVec
	readTimeGauge *prom.GaugeVec
	conf          config.Config
	topics        []string
	mu            sync.Mutex
	rnd           *rand.Rand
)

func init() {
	conf = config.GetConfig()
	topicsGauge = prom.NewGaugeVec(prom.GaugeOpts{Name: "topics"}, []string{"type"})
	readTimeGauge = prom.NewGaugeVec(prom.GaugeOpts{Name: "read_time"}, []string{"type"})
	prom.MustRegister(topicsGauge)
	prom.MustRegister(readTimeGauge)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func CountMetrics() {
	c := client.NewHTTPClient(conf.ConsumerKey, conf.AccessToken, conf.Host, conf.Path)
	for {
		result, err := c.GetTopicsList()
		if err != nil {
			log.Error(err.Error())
			continue
		}
		var (
			unreadCount   int64
			archivedCount int64
			unreadTime    int64
			archivedTime  int64
		)
		tmpTopics := make([]string, 0, len(result.List))
		for _, topic := range result.List {
			switch topic.Status {
			case domain.UNREAD_STATUS:
				unreadCount++
				unreadTime += topic.TimeToRead
			case domain.ARCHIVED_STATUS:
				archivedCount++
				archivedTime += topic.TimeToRead
			}
			tmpTopics = append(tmpTopics, topic.ItemID)
		}
		mu.Lock()
		topics = tmpTopics
		mu.Unlock()
		log.Info("unread: %d, unread_time: %d, archived: %d, archived_time: %d", unreadCount, unreadTime, archivedCount, archivedTime)
		topicsGauge.WithLabelValues("unread").Set(float64(unreadCount))
		readTimeGauge.WithLabelValues("unread").Set(float64(unreadTime))
		topicsGauge.WithLabelValues("archived").Set(float64(archivedCount))
		readTimeGauge.WithLabelValues("archived").Set(float64(archivedTime))
	}
}

func GetRandomTopicID() string {
	mu.Lock()
	defer mu.Unlock()
	return topics[rnd.Int63n(int64(len(topics)))]
}
