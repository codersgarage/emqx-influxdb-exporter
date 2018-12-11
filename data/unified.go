package data

import (
	"errors"
	"fmt"
	"github.com/codersgarage/emqx-influxdb-exporter/env"
	"github.com/codersgarage/emqx-influxdb-exporter/log"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
	"github.com/nahid/gohttp"
)

type UnifiedDao struct {
}

func NewUnifiedDao() *UnifiedDao {
	return &UnifiedDao{}
}

func (ud *UnifiedDao) Write(m models.Unified) error {
	t := ""
	for k, v := range m.Payload {
		t = fmt.Sprintf(`%s
%s %s=%di`, t, m.CollectionName, k, v)
	}

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
