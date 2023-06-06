package services

type ReservationServiceMock struct {
}

func (s ReservationServiceMock) CreateReservation(hotelID int, checkin, checkout, token, clientName string) error {

	return nil
}
