package infra

import (
	"errors"

	validator "github.com/asaskevich/govalidator"
)

type PositionValidator interface {
	Validate(latitude, longitude float64) error
}

type PositionGoValidator struct{}

func (v *PositionGoValidator) Validate(latitude, longitude float64) error {
	validLng := validator.IsLongitude(validator.ToString(longitude))
	validLat := validator.IsLatitude(validator.ToString(latitude))

	if !validLat || !validLng {
		return errors.New("invalid latitude or longitude format")
	}

	return nil
}
