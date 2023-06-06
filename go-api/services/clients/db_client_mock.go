package clients

import (
	"fmt"
	"go-api/dao"
)

type DBClientMock struct{}

func (DBClientMock) CreateReservation(reservation dao.Reservation) error {
	fmt.Println(fmt.Sprintf("Creating reservation: %v", reservation))
	return nil
}

func (DBClientMock) GetHotelByID(hotelID int) (dao.Hotel, error) {
	fmt.Println(fmt.Sprintf("Getting hotel by ID: %d", hotelID))
	return dao.Hotel{
		ID:          1,
		Name:        "Test Hotel",
		Photo:       "Test Photo.jpg",
		Description: "Test Description",
		Location:    "Test Location",
		Rooms:       10,
	}, nil
}

func (DBClientMock) CountReservations(hotelID int, checkin string, checkout string) (int, error) {
	fmt.Println(fmt.Sprintf("Counting reservations for hotel: %d", hotelID))
	return 5, nil
}
