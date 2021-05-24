package controllers

import (
	"gps-worker/domain"
	"gps-worker/usecases"
)

type Controller struct {
	usecases.PositionService
}

func (c *Controller) CreateService(latitude, longitude float64) (position domain.Position, err error) {
	position, err = c.Create(latitude, longitude)
	return
}
