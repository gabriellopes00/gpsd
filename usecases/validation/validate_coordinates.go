package validation

import (
	"errors"
	"fmt"
	"github.com/gabriellopes00/gpsd/domain"
	"regexp"
)

type CoordinatesValidator struct {}

func (v *CoordinatesValidator) Validate(positions ...domain.Position) error {

	for _, position := range positions {
		if !v.isValidLat(position.Latitude) {
			return errors.New("invalid longitude format")

		} else if !v.isValidLng(position.Longitude) {
			return errors.New("invalid latitude format")

		}
	}

	return nil
}

func (v *CoordinatesValidator) isValidLat(lat float64) bool {
	const pattern string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	rx := regexp.MustCompile(pattern)

	latStr := fmt.Sprintf("%v", lat)

	return rx.MatchString(latStr)
}

func (v *CoordinatesValidator) isValidLng(lng float64) bool {
	const pattern string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	rx := regexp.MustCompile(pattern)

	lngStr := fmt.Sprintf("%v", lng)

	return rx.MatchString(lngStr)
}
