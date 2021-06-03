package calc

import (
	d "gps-worker/domain"
)

func GetDistances(entrypoint *d.Position, positions []*d.Position) {
	for _, p := range positions {
		p.Distance = CalcDistance(
			entrypoint.Latitude,
			entrypoint.Longitude,
			p.Latitude,
			p.Longitude,
		)
	}
}
