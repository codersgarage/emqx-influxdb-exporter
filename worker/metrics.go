package worker

import (
	"fmt"
	"github.com/codersgarage/emqx-influxdb-exporter/data"
	"github.com/codersgarage/emqx-influxdb-exporter/env"
	"github.com/codersgarage/emqx-influxdb-exporter/log"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
	"github.com/nahid/gohttp"
	"time"
)

func RunMetricsWorker() {
	repo := data.NewMetricsDao()

	for {
		resp, err := gohttp.NewRequest().
			BasicAuth(env.GetEmqxAppID(), env.GetEmqxAppSecret()).
			Get(fmt.Sprintf("%s/api/v3/metrics", env.GetEmqxURL()))

		if err != nil {
			log.Log().Errorln("Failed - ", err)
			return
		}

		var res []models.Metrics
		if err := resp.UnmarshalBody(&res); err != nil {
			log.Log().Errorln("Failed - ", err)
			return
		}

		met := models.Metrics{}
		for _, m := range res {
			v := models.Vol{}
			v.BytesReceived += m.Metrics.BytesReceived
			v.BytesSent += m.Metrics.BytesSent
			v.MessagesExpired += m.Metrics.MessagesExpired
			v.PacketsPublishReceived += m.Metrics.PacketsPublishReceived
			v.PacketsPublishSent += m.Metrics.PacketsPublishSent
			v.MessagesQos1Received += m.Metrics.MessagesQos1Received
			v.MessagesQos1Sent += m.Metrics.MessagesQos1Sent
			v.MessagesReceived += m.Metrics.MessagesReceived
			v.MessagesSent += m.Metrics.MessagesSent
			v.MessagesDropped += m.Metrics.MessagesDropped
			v.MessagesForward += m.Metrics.MessagesForward

			met.Metrics = v

			m.SetCollectionName(fmt.Sprintf("metrics_%s", m.Node))
			if err := repo.Write(m); err != nil {
				log.Log().Errorln("Write failed - ", err)
				return
			}
			log.Log().Infoln("Write success - ", m.CollectionName())
		}

		met.SetCollectionName("metrics")
		if err := repo.Write(met); err != nil {
			log.Log().Errorln("Write failed - ", err)
			return
		}
		log.Log().Infoln("Write success - ", met.CollectionName())

		time.Sleep(5 * time.Second)
	}
}
