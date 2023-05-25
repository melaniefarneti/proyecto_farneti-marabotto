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
	baseURL       = "https://example.com/api%s" //VER ESTO
)

func (client HTTPClient) GetHotelByID(id int) (Hotel, error) {
	// Build the URL
	endpoint := fmt.Sprintf(hotelEndpoint, id)
	url := fmt.Sprintf(baseURL, endpoint)

	// Invoke API
	response, err := http.Get(url)
	if err != nil {
		return Hotel{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return Hotel{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	// Read response payload bytes
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Hotel{}, err
	}

	// Convert bytes to custom struct
	var hotel Hotel
	err = json.Unmarshal(bytes, &hotel)
	if err != nil {
		return Hotel{}, err
	}

	return hotel, nil
}
