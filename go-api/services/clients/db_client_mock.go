package clients

import (
	"fmt"
	"go-api/dao"
)

type DBClientMock struct {
	//Hotels      []dao.Hotel
	//DeletedIDs  []int
	//DeleteError error
}

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

func (DBClientMock) GetHotels() ([]dao.Hotel, error) {
	fmt.Println("Getting hotels")
	hotels := []dao.Hotel{
		{
			ID:          1,
			Name:        "Hotel 1",
			Photo:       "photo1.jpg",
			Description: "Description 1",
			Location:    "Location 1",
			Rooms:       10,
		},
		{
			ID:          2,
			Name:        "Hotel 2",
			Photo:       "photo2.jpg",
			Description: "Description 2",
			Location:    "Location 2",
			Rooms:       20,
		},
	}
	return hotels, nil
}

/*
func (m DBClientMock) CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error) {
	fmt.Println(fmt.Sprintf("Creating hotel: %v", hotel))
	return hotel, nil
}

func (m DBClientMock) DeleteHotel(hotelID int) error {
	m.DeletedIDs = append(m.DeletedIDs, hotelID)
	return m.DeleteError
}
*/