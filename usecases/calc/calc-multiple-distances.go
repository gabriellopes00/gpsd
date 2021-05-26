package calc

import (
	d "gps-worker/domain"
	i "gps-worker/infra"
	"log"
	m "math"
)

func CalcMultipleDistances(entrypoint *d.Position, positions []d.Position) []d.Position {
	var calculated []d.Position
	for _, p := range positions {
		err := i.Validate(p.Latitude, p.Longitude)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		result := CalcDistance(entrypoint.Latitude, entrypoint.Longitude, p.Latitude, p.Longitude)
		p.Distance = result
		calculated = append(calculated, p)
	}

	return calculated
}

// Round float values to integer ones
func round(num float64) int {
	return int(num + m.Copysign(0.5, num))
}

// Truncate distance decimal cases
func toFixed(num float64, precision int) float64 {
	output := m.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
