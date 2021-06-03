package validation

import (
	d "gps-worker/domain"
	i "gps-worker/infra"
)

func ValidateCoordinates(positions []*d.Position) error {
	for _, p := range positions {
		if err := i.Validate(p.Latitude, p.Longitude); err != nil {
			return err
		}
	}
	return nil
}
