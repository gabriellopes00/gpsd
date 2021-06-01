package calc

import (
	d "gps-worker/domain"
)

func GetRadius(entrypoint d.Position, positions []d.Position) []d.Position {
	LatRadius := []float64{entrypoint.Latitude + 0.1, entrypoint.Latitude - 0.1}
	LngRadius := []float64{entrypoint.Longitude + 0.1, entrypoint.Longitude - 0.1}

	var radius []d.Position
	for _, p := range positions {
		if (p.Latitude >= LatRadius[1] && p.Latitude <= LatRadius[0]) &&
			(p.Longitude >= LngRadius[1] && p.Longitude <= LngRadius[0]) {
			radius = append(radius, p)
		}
	}

	return radius
}
