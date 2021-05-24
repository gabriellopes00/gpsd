package domain_test

import (
	"gps-worker/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPosition(t *testing.T) {
	_, err := domain.NewPosition(-24.00097794064844, -46.1797712784327)
	require.Nil(t, err)

	_, err = domain.NewPosition(-890.00, -46.1797712784327)
	require.Error(t, err)

	_, err = domain.NewPosition(-24.00097794064844, -460.17)
	require.Error(t, err)
}
