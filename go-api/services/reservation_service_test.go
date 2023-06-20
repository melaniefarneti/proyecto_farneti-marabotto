package services

import (
	"errors"
	"go-api/clients"
	"go-api/dao"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	// Setup
	dbClientMock := &clients.DBClientMock{}
	reservationServiceMock := &ReservationServiceMock{
		DBClientMock: dbClientMock,
	}

	// Test Case 1: Successful reservation creation
	reservation := dao.Reservation{
		HotelID:    1,
		CheckIn:    time.Now().Format("2006-01-02"),
		CheckOut:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
		ClientName: "John Doe",
	}

	err := reservationServiceMock.CreateReservation(reservation.HotelID, reservation.CheckIn, reservation.CheckOut, reservation.ClientName)
	assert.NoError(t, err, "Error should be nil")

	// Test Case 2: Error in reservation creation
	reservationServiceMock.CreateReservationFunc = func(reservation dao.Reservation) error {
		return errors.New("reservation creation error")
	}

	err = reservationServiceMock.CreateReservation(reservation.HotelID, reservation.CheckIn, reservation.CheckOut, reservation.ClientName)
	assert.Error(t, err, "Error should not be nil")
}
