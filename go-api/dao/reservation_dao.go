package dao

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	HotelID    int `gorm:"column:hotel_id"`
	Hotel      Hotel
	CheckIn    string `gorm:"column:fecha_desde"`
	CheckOut   string `gorm:"column:fecha_hasta"`
	ClientName string `gorm:"column:cliente_nombre"`
}
