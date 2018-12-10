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

func RunStatWorker() {
	repo := data.NewStatsDao()

	for {
		resp, err := gohttp.NewRequest().
			BasicAuth(env.GetEmqxAppID(), env.GetEmqxAppSecret()).
			Get(fmt.Sprintf("%s/api/v3/stats", env.GetEmqxURL()))

		if err != nil {
			log.Log().Errorln("Failed - ", err)
			return
		}

		var res []models.Stats
		if err := resp.UnmarshalBody(&res); err != nil {
			log.Log().Errorln("Failed - ", err)
			return
		}

		stat := models.Stats{}
		for _, s := range res {
			stat.ConnectionsCount += s.ConnectionsCount
			stat.ConnectionsMax += s.ConnectionsMax
			stat.TopicsCount += s.TopicsCount
			stat.TopicsMax += s.TopicsMax
			stat.SubscribersCount += s.SubscribersCount
			stat.SubscribersMax += s.SubscribersMax
			stat.SubscriptionsCount += s.SubscriptionsCount
			stat.SubscriptionsMax += s.SubscriptionsMax
			stat.RoutesCount += s.RoutesCount
			stat.RoutesMax += s.RoutesMax

			s.SetCollectionName(fmt.Sprintf("stats_%s", s.Node))
			if err := repo.Write(s); err != nil {
				log.Log().Errorln("Write failed - ", err)
				return
			}
			log.Log().Infoln("Write success")
		}

		stat.SetCollectionName("stats")
		if err := repo.Write(stat); err != nil {
			log.Log().Errorln("Write failed - ", err)
			return
		}
		log.Log().Infoln("Write success - ", stat.CollectionName())

		time.Sleep(5 * time.Second)
	}
}
