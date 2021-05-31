package calc

import (
	"fmt"
	"gps-worker/domain"
	"math"
	"strconv"
)

func truncate(value float64) float64 {
	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Ceil(value*100)/100), 64)
	return result
}

func GetRadius(entrypoint domain.Position, positions []domain.Position) []domain.Position {
	lat := truncate(entrypoint.Latitude)
	lng := truncate(entrypoint.Longitude)

	LatRadius := []float64{lat + 0.1, lat - 0.1}
	LngRadius := []float64{lng + 0.1, lng - 0.1}

	fmt.Println(LatRadius, LngRadius)

	var radius []domain.Position
	for _, p := range positions {
		p.Latitude = truncate(p.Latitude)
		p.Longitude = truncate(p.Longitude)

		if (p.Latitude >= LatRadius[1] && p.Latitude <= LatRadius[0]) &&
			(p.Longitude >= LngRadius[1] && p.Longitude <= LngRadius[0]) {
			radius = append(radius, p)

			if p.Name == "SP" {
				fmt.Println(p)
			}
		}
	}

	// for _, v := range radius {
	// 	fmt.Println(v)
	// }

	return radius
}
