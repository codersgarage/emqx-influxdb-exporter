package api

import (
	"github.com/codersgarage/emqx-influxdb-exporter/app"
	"github.com/codersgarage/emqx-influxdb-exporter/repos"
	"net/http"
)

type UnifiedRepo interface {
	UnifiedStat(s *app.Scope) error
}

type UnifiedRoute struct {
	Repo UnifiedRepo
}

func NewUnifiedRoute() *UnifiedRoute {
	return &UnifiedRoute{
		Repo: repos.NewUnifiedRepo(),
	}
}

func (sr *UnifiedRoute) unifiedStat(w http.ResponseWriter, r *http.Request) {
	resp := response{}

	if err := sr.Repo.UnifiedStat(app.NewScope(r)); err != nil {
		resp.Title = "Failed"
		resp.Status = http.StatusInternalServerError
		resp.Errors = err
		resp.ServerJSON(w)
		return
	}

	resp.Title = "Success"
	resp.Status = http.StatusOK
	resp.ServerJSON(w)
}
