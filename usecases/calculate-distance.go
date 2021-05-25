package usecases

import (
	"fmt"
	d "gps-worker/domain"
	i "gps-worker/infra"
	"log"
	m "math"
)

func CalcDistance(entrypoint *d.Position, positions []d.Position) {
	var distances []float64
	var calculatedPositions []d.Position

	for _, value := range positions {
		err := i.Validate(value.Latitude, value.Longitude)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		result := calc(entrypoint.Latitude, entrypoint.Longitude, value.Latitude, value.Longitude)
		distances = append(distances, result)
		calculatedPositions = append(calculatedPositions, value)
	}

	for i1 := range distances {
		for i2 := range distances {
			if distances[i1] < distances[i2] {
				distances[i1], distances[i2] = distances[i2], distances[i1]
				calculatedPositions[i1], calculatedPositions[i2] = calculatedPositions[i2], calculatedPositions[i1]
			}
		}
	}

	for _, v := range calculatedPositions {
		fmt.Println(v)
	}
}

// Calculate distance between two coordinates
func calc(lat1, lng1, lat2, lng2 float64) float64 {
	earthRadiusKm := 6371
	dLat := deg2rad(lat2 - lat1)
	dLng := deg2rad(lng2 - lng1)

	lat1 = deg2rad(lat1)
	lat2 = deg2rad(lat2)

	a := m.Sin(dLat/2)*m.Sin(dLat/2) + m.Sin(dLng/2)*m.Sin(dLng/2)*m.Cos(lat1)*m.Cos(lat2)
	var c = 2 * m.Atan2(m.Sqrt(a), m.Sqrt(1-a))

	result := float64(earthRadiusKm) * float64(c)
	return toFixed((result * 1000), 1)
}

// Convert coordinates degrees to radians
func deg2rad(deg float64) float64 {
	return (deg * m.Pi) / 180
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
