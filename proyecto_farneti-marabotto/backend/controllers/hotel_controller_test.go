package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHotelParams(t *testing.T) {
	assert.Equal(t, "hotelID", paramHotelID)
}
