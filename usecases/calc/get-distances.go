package calc

import (
	d "gps-worker/domain"
)

func GetDistances(entrypoint *d.Position, positions []d.Position) []d.Position {
	var distances []d.Position
	for _, p := range positions {
		distance := CalcDistance(entrypoint.Latitude, entrypoint.Longitude, p.Latitude, p.Longitude)
		p.Distance = distance
		distances = append(distances, p)
	}
	return distances
}
