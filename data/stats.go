package data

import (
	"errors"
	"fmt"
	"github.com/codersgarage/emqx-influxdb-exporter/env"
	"github.com/codersgarage/emqx-influxdb-exporter/log"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
	"github.com/nahid/gohttp"
)

type StatsDao struct {
}

func NewStatsDao() *StatsDao {
	return &StatsDao{}
}

func (sd *StatsDao) Write(s models.Stats) error {
	t := fmt.Sprintf(`%s node="%s"`, s.CollectionName(), s.Node)
	t = fmt.Sprintf(`%s
%s subscriptions_count=%di`, t, s.CollectionName(), s.SubscriptionsCount)
	t = fmt.Sprintf(`%s
%s subscriptions_max=%di`, t, s.CollectionName(), s.SubscriptionsMax)
	t = fmt.Sprintf(`%s
%s topics_count=%di`, t, s.CollectionName(), s.TopicsCount)
	t = fmt.Sprintf(`%s
%s topics_max=%di`, t, s.CollectionName(), s.TopicsMax)
	t = fmt.Sprintf(`%s
%s routes_count=%di`, t, s.CollectionName(), s.RoutesCount)
	t = fmt.Sprintf(`%s
%s routes_max=%di`, t, s.CollectionName(), s.RoutesMax)
	t = fmt.Sprintf(`%s
%s subscribers_count=%di`, t, s.CollectionName(), s.SubscribersCount)
	t = fmt.Sprintf(`%s
%s subscribers_max=%di`, t, s.CollectionName(), s.SubscribersMax)
	t = fmt.Sprintf(`%s
%s connections_count=%di`, t, s.CollectionName(), s.ConnectionsCount)
	t = fmt.Sprintf(`%s
%s connections_max=%di`, t, s.CollectionName(), s.ConnectionsMax)

	log.Log().Infoln(t)

	resp, err := gohttp.NewRequest().
		Body([]byte(t)).
		Post(env.GetInfluxURL())

	if err != nil {
		return err
	}

	if resp.GetStatusCode() >= 200 && resp.GetStatusCode() < 300 {
		return nil
	}
	return errors.New(fmt.Sprintf("Failed with status : %d", resp.GetStatusCode()))
}
