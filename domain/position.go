package domain

import (
	"errors"

	validator "github.com/asaskevich/govalidator"
)

type Position struct {
	Latitude  float64 `json:"latitude,omitempty" valid:"latitude"`
	Longitude float64 `json:"longitude,omitempty" valid:"longitude"`
}

func NewPosition(latitude, longitude float64) (*Position, error) {
	newPosition := Position{Latitude: latitude, Longitude: longitude}

	validator.SetFieldsRequiredByDefault(true)
	_, err := validator.ValidateStruct(newPosition)
	if err != nil {
		return nil, errors.New("invalid latitude or longitude format")
	}

	return &newPosition, nil
}
