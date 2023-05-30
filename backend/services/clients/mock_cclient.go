package clients

// MockClient es una implementación particular de HotelClient (interfaz)
// Funciona devolviendo valores predefinidos
// Este cliente no necesita ser probado :)
type MockClient struct{}

func (mockClient MockClient) GetHotelByID(id int) (Hotel, error) {
	return Hotel{
		ID:    1,
		Name:  "Mocked Hotel",
		Price: 100.00,
	}, nil
}
