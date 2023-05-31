package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReservationParams(t *testing.T) {
	assert.Equal(t, "reservationID", paramReservationID)
}
