package calc

import (
	d "gps-worker/domain"
)

func Filter(entrypoint *d.Position, positions *[]d.Position) (*[]d.Position, error) {
	return positions, nil
}
