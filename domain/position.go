package domain

import "time"

type Position struct {
	Latitude  float64   `json:"latitude,omitempty" valid:"latitude"`
	Longitude float64   `json:"longitude,omitempty" valid:"longitude"`
	CreatedAt time.Time `json:"created_at,omitempty" valid:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" valid:"-"`
}

// package domain_test

// import (
// 	"gps-worker/domain"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestNewPosition(t *testing.T) {
// 	_, err := domain.NewPosition(-24.00097794064844, -46.1797712784327)
// 	require.Nil(t, err)

// 	_, err = domain.NewPosition(-890.00, -46.1797712784327)
// 	require.Error(t, err)

// 	_, err = domain.NewPosition(-24.00097794064844, -460.17)
// 	require.Error(t, err)
// }

// func NewPosition(latitude, longitude float64) (*Position, error) {
// 	newPosition := Position{Latitude: latitude, Longitude: longitude}

// 	validator.SetFieldsRequiredByDefault(true)
// 	_, err := validator.ValidateStruct(newPosition)
// 	if err != nil {
// 		return nil, errors.New("invalid latitude or longitude format")
// 	}

// 	return &newPosition, nil
// }
