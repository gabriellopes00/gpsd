package calc

import d "gps-worker/domain"

func Filter(positions []*d.Position) []*d.Position {
	if len(positions) <= 5 {
		return positions
	} else {
		return positions[:5]
	}
}
