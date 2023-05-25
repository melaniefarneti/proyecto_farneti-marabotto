package clients

// Hotel is a struct that represents a hotel
type Hotel struct {
	ID    int
	Name  string
	Price float64
}

// HotelClient is the interface definition for a hotel client
// It only defines the contract
type HotelClient interface {
	GetHotelByID(id int) (Hotel, error)
}

