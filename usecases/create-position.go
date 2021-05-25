package usecases

import (
	"gps-worker/domain"
	"gps-worker/infra"
	"time"
)

type PositionService interface {
	Create(latitude, longitude float64) (position domain.Position, err error)
}

type Service struct {
	infra.FakeRepository
	infra.PositionGoValidator
}

func (service *Service) Create(latitude, longitude float64) (position domain.Position, err error) {
	err = service.Validate(latitude, longitude)
	if err != nil {
		return position, err
	}

	position = domain.Position{
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	service.Insert(position)
	return
}
