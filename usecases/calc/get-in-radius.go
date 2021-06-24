package calc

import d "gps-worker/domain"

func GetInRadius(entrypoint *d.Position, positions []*d.Position) []*d.Position {
	LatRange := []float64{entrypoint.Latitude + 0.1, entrypoint.Latitude - 0.1}
	LngRange := []float64{entrypoint.Longitude + 0.1, entrypoint.Longitude - 0.1}

	var radius []*d.Position
	for _, p := range positions {
		if (p.Latitude >= LatRange[1] && p.Latitude <= LatRange[0]) &&
			(p.Longitude >= LngRange[1] && p.Longitude <= LngRange[0]) {
			radius = append(radius, p)
		}
	}

	return radius
}
