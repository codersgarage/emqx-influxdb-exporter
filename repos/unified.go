package repos

import (
	"github.com/codersgarage/emqx-influxdb-exporter/app"
	"github.com/codersgarage/emqx-influxdb-exporter/data"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
)

type UnifiedRepo struct {
	Dao *data.UnifiedDao
}

func NewUnifiedRepo() *UnifiedRepo {
	return &UnifiedRepo{
		Dao: data.NewUnifiedDao(),
	}
}

func (ur *UnifiedRepo) UnifiedStat(s *app.Scope) error {
	u := models.Unified{}
	if err := u.Parse(s); err != nil {
		return err
	}

	if err := ur.Dao.Write(u); err != nil {
		return err
	}
	return nil
}
