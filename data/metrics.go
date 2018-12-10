package data

import (
	"errors"
	"fmt"
	"github.com/codersgarage/emqx-influxdb-exporter/env"
	"github.com/codersgarage/emqx-influxdb-exporter/log"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
	"github.com/nahid/gohttp"
)

type MetricsDao struct {
}

func NewMetricsDao() *MetricsDao {
	return &MetricsDao{}
}

func (md *MetricsDao) Write(m models.Metrics) error {
	t := fmt.Sprintf(`%s node="%s"`, m.CollectionName(), m.Node)
	t = fmt.Sprintf(`%s
%s bytes_received=%di`, t, m.CollectionName(), m.Metrics.BytesReceived)
	t = fmt.Sprintf(`%s
%s bytes_sent=%di`, t, m.CollectionName(), m.Metrics.BytesSent)
	t = fmt.Sprintf(`%s
%s messages_expired=%di`, t, m.CollectionName(), m.Metrics.MessagesExpired)
	t = fmt.Sprintf(`%s
%s packets_publish_received=%di`, t, m.CollectionName(), m.Metrics.PacketsPublishReceived)
	t = fmt.Sprintf(`%s
%s packets_publish_sent=%di`, t, m.CollectionName(), m.Metrics.PacketsPublishSent)
	t = fmt.Sprintf(`%s
%s messages_qos1_received=%di`, t, m.CollectionName(), m.Metrics.MessagesQos1Received)
	t = fmt.Sprintf(`%s
%s messages_qos1_sent=%di`, t, m.CollectionName(), m.Metrics.MessagesQos1Sent)
	t = fmt.Sprintf(`%s
%s messages_received=%di`, t, m.CollectionName(), m.Metrics.MessagesReceived)
	t = fmt.Sprintf(`%s
%s messages_sent=%di`, t, m.CollectionName(), m.Metrics.MessagesSent)
	t = fmt.Sprintf(`%s
%s messages_dropped=%di`, t, m.CollectionName(), m.Metrics.MessagesDropped)
	t = fmt.Sprintf(`%s
%s messages_forward=%di`, t, m.CollectionName(), m.Metrics.MessagesForward)

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
