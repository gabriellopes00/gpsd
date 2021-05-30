package calc

import (
	"errors"
	d "gps-worker/domain"
)

func Filter(positions []d.Position, offset, skip uint8) ([]d.Position, error) {
	err := errors.New("invalid offset and skip values")

	if (skip+offset >= uint8(len(positions))) || offset <= 0 {
		return nil, err
	}

	return positions[skip:offset], nil
}
