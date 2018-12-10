package repos

import (
	"github.com/codersgarage/emqx-influxdb-exporter/app"
	"github.com/codersgarage/emqx-influxdb-exporter/data"
	"github.com/codersgarage/emqx-influxdb-exporter/errors"
	"github.com/codersgarage/emqx-influxdb-exporter/models"
	"github.com/codersgarage/emqx-influxdb-exporter/utils"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

type MonkeyRepo struct {
	Dao *data.MonkeyDao
}

func NewMonkeyRepo() *MonkeyRepo {
	return &MonkeyRepo{
		Dao: data.NewMonkeyDao(),
	}
}

func (hr *MonkeyRepo) SaveMonkey(s *app.Scope) (*models.Monkey, error) {
	monkey := &models.Monkey{}
	if err := utils.ParseBody(s.Request, monkey); err != nil {
		return nil, errors.NewAPIError(http.StatusBadRequest, "400001", "Unable to parse request body", nil)
	}
	if err := monkey.Validate(); err != nil {
		return nil, errors.NewAPIError(http.StatusUnprocessableEntity, "422001", "Invalid data", err)
	}

	monkey.CreatedAt = time.Now()
	monkey.UpdatedAt = time.Now()

	if err := hr.Dao.SaveMonkey(s, monkey); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to save monkey", err)
	}
	return monkey, nil
}

func (hr *MonkeyRepo) GetMonkey(s *app.Scope) (*models.Monkey, error) {
	monkey := &models.Monkey{}

	ID := chi.URLParam(s.Request, "id")
	err := hr.Dao.GetMonkey(s, ID, monkey)

	if err != nil && errors.IsRecordNotFoundError(err) {
		return nil, errors.NewAPIError(http.StatusNotFound, "404001", "monkey not found", err)
	}
	if err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to get monkey", err)
	}
	return monkey, nil
}

func (hr *MonkeyRepo) UpdateMonkey(s *app.Scope) (*models.Monkey, error) {
	return &models.Monkey{}, nil
}

func (hr *MonkeyRepo) ListMonkey(s *app.Scope) ([]models.Monkey, error) {
	var data []models.Monkey
	if err := hr.Dao.ListMonkeys(s, &data); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to list Monkeys", err)
	}
	if len(data) > 0 {
		return data, nil
	}
	return []models.Monkey{}, nil
}

func (hr *MonkeyRepo) DeleteMonkey(s *app.Scope) error {
	return nil
}
