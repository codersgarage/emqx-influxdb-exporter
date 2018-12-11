package models

import (
	"github.com/codersgarage/emqx-influxdb-exporter/app"
	"github.com/codersgarage/emqx-influxdb-exporter/utils"
)

type Unified struct {
	CollectionName string           `json:"collection_name"`
	Payload        map[string]int64 `json:"payload"`
}

func (u *Unified) Parse(s *app.Scope) error {
	return utils.ParseBody(s.Request, u)
}
