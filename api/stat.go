package api

import "net/http"

type StatRepo interface {
	SendStat() error
}

type StatRoute struct {
	Repo StatRepo
}

func NewStatRoute() *StatRoute {
	return &StatRoute{}
}

func (sr *StatRoute) sendStat(w http.ResponseWriter, r *http.Request) {

}
