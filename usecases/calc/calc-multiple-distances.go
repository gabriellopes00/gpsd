package calc

import (
	d "gps-worker/domain"
	i "gps-worker/infra"
	"log"
)

func CalcMultipleDistances(entrypoint *d.Position, positions []d.Position) []d.Position {
	var calculated []d.Position
	for _, p := range positions {
		err := i.Validate(p.Latitude, p.Longitude)
		if err != nil {
			log.Fatalln(err)
			break
		}

		distance := CalcDistance(entrypoint.Latitude, entrypoint.Longitude, p.Latitude, p.Longitude)
		p.Distance = distance
		calculated = append(calculated, p)
	}

	return calculated
}
