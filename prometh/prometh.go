package prometh

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
)

func Push(count float64, url string) {
	job := "filecoin_chain_message_count_per_height"
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{Name: job})
	gauge.Set(count)
	err := push.New(url, job).Grouping("module", "chain").Collector(gauge).Push()
	if err != nil {
		log.Printf("push prometheus %s failed:%s", url, err)
	}
}
