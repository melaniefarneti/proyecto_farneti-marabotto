package services

import (
	"go-api/clients"
	"go-api/dao"
)

type ReservationServiceMock struct {
	DBClientInterface     clients.DBClientInterface
	DBClientMock          *clients.DBClientMock
	CreateReservationFunc func(reservation dao.Reservation) error
	GetReservationsFunc   func() ([]dao.Reservation, error)
	CountReservationsFunc func(hotelID int, checkin string, checkout string) (int, error)
	GetHotelByIDFunc      func(hotelID int) (dao.Hotel, error)
}

func (m *ReservationServiceMock) CreateReservation(hotelID int, checkin, checkout, clientName string) error {
	if m.CreateReservationFunc != nil {
		reservation := dao.Reservation{
			HotelID:    hotelID,
			CheckIn:    checkin,
			CheckOut:   checkout,
			ClientName: clientName,
		}
		return m.CreateReservationFunc(reservation)
	}
	return nil
}

func (m *ReservationServiceMock) GetReservations() ([]dao.Reservation, error) {
	if m.GetReservationsFunc != nil {
		return m.GetReservationsFunc()
	}
	return nil, nil
}

func (m *ReservationServiceMock) CountReservations(hotelID int, checkin string, checkout string) (int, error) {
	if m.CountReservationsFunc != nil {
		return m.CountReservationsFunc(hotelID, checkin, checkout)
	}
	return 0, nil
}

func (m *ReservationServiceMock) GetHotelByID(hotelID int) (dao.Hotel, error) {
	if m.GetHotelByIDFunc != nil {
		return m.GetHotelByIDFunc(hotelID)
	}
	return dao.Hotel{}, nil
}
