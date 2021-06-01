package validation

import (
	d "gps-worker/domain"
	i "gps-worker/infra"
	"log"
)

func ValidateCoordinates(positions []d.Position) error {
	for _, p := range positions {
		err := i.Validate(p.Latitude, p.Longitude)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}

	return nil
}
