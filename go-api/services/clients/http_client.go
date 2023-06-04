package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPClient struct{}

const (
	hotelEndpoint = "/hotels/%d"
	baseURL       = "https://api.reservahotelera.com/api" //ME LO INVENTE VER QUE LINK VA ACA
)

func (client HTTPClient) GetHotels() ([]MLHotel, error) {
	// Construir la URL
	url := fmt.Sprintf(baseURL + "/hotels")

	// Invocar a la API
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Validar el error de la API
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("código de estado inesperado %d", response.StatusCode))
	}

	// Leer los bytes de la respuesta
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Convertir los bytes a una lista de estructuras de hoteles
	var hotels []MLHotel
	err = json.Unmarshal(bytes, &hotels)
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

/*func (client HTTPClient) GetHotelByID(id int) (MLHotel, error) {
	// Construir la URL
	url := fmt.Sprintf(baseURL+hotelEndpoint, id)

	// Invocar a la API
	response, err := http.Get(url)
	if err != nil {
		return MLHotel{}, err
	}

	// Validar el error de la API
	if response.StatusCode != http.StatusOK {
		return MLHotel{}, errors.New(fmt.Sprintf("código de estado inesperado %d", response.StatusCode))
	}

	// Leer los bytes de la respuesta
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return MLHotel{}, err
	}

	// Convertir los bytes a una estructura personalizada
	var hotel MLHotel
	err = json.Unmarshal(bytes, &hotel)
	if err != nil {
		return MLHotel{}, err
	}

	return hotel, nil
}*/
