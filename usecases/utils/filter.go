package utils

import (
	"errors"
	"github.com/gabriellopes00/gpsd/domain"
)

func Filter(positions []*domain.Position, skip int, offset int) ([]*domain.Position, error) {
	if len(positions) <= skip + offset {
		return nil, errors.New("filter params out of range")
	}

	return positions[skip:offset], nil
}
