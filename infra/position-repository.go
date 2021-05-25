package infra

import "gps-worker/domain"

type PositionRepository interface {
	Insert(data domain.Position) (result domain.Position, err error)
}

type FakeRepository struct{}

func (r *FakeRepository) Insert(data domain.Position) (result domain.Position, err error) {
	result = domain.Position{
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return
}
