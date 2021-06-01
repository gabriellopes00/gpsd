package calc

import (
	d "gps-worker/domain"
)

func GetDistances(entrypoint *d.Position, channel chan d.Position) {
	for p := range channel {
		distance := CalcDistance(entrypoint.Latitude, entrypoint.Longitude, p.Latitude, p.Longitude)
		p.Distance = distance
		channel <- p
	}
}
