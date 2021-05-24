package infra

import (
	"errors"

	validator "github.com/asaskevich/govalidator"
)

type PositionValidator interface {
	Validate(latitude, longitude float64) error
}

func Validate(latitude, longitude float64) error {
	validLng := validator.IsLongitude(validator.ToString(longitude))
	validLat := validator.IsLatitude(validator.ToString(latitude))

	if !validLat || !validLng {
		return errors.New("invalid latitude or longitude format")
	}

	return nil
}
