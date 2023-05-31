package clients

type MLHotel struct {
	ID   int
	Name string
}

type MLClient interface {
	GetHotels() ([]MLHotel, error)
}
