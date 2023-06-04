package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
	assert.Equal(t, "/hotels/:hotelID", pathGetHotel)       // Verifica si el valor de PathGetHotel es igual a "/hotels/:hotelID".
	assert.Equal(t, "/reservations", pathCreateReservation) // Verifica si el valor de PathCreateReservation es igual a "/reservations".
	//assert.Equal(t, "/amenities/:amenityID", PathGetAmenity) // Verifica si el valor de PathGetAmenity es igual a "/amenities/:amenityID".
	assert.Equal(t, "/users/:userID", pathGetUser) // Verifica si el valor de PathGetUser es igual a "/users/:userID".
}
