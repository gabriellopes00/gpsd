package calc

import "github.com/gabriellopes00/gpsd/domain"

func GetDistances(entrypoint *domain.Position, positions []*domain.Position) {
	for _, p := range positions {
		p.Distance = Distance(entrypoint.Latitude, entrypoint.Longitude, p.Latitude, p.Longitude)
	}
}
