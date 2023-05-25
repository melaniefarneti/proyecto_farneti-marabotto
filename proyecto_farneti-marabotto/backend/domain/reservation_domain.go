package domain

type Reservation struct {
	ID         int    `gorm:"primaryKey"`
	HotelID    int    `gorm:"column:hotel_id"`
	CheckIn    string `gorm:"column:fecha_desde"`
	CheckOut   string `gorm:"column:fecha_hasta"`
	ClientName string `gorm:"column:cliente_nombre"`
}
