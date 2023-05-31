package clients

import (
	"go-api/domain"
)

// MockClient es una implementación de la interfaz MLClient para pruebas simuladas.
type MockClient struct{}

// GetHotels devuelve un listado de hoteles desde una fuente de datos simulada.
func (c MockClient) GetHotels() ([]domain.Hotel, error) {
	// Simulamos una lista de hoteles
	fakeHotels := []domain.Hotel{
		{ID: 1, Name: "Hotel A"},
		{ID: 2, Name: "Hotel B"},
		{ID: 3, Name: "Hotel C"},
	}
	return fakeHotels, nil
}
