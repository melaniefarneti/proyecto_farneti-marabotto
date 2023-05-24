package domain
//ESTO DEBEMOS CAMBIARLO

type Item struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}