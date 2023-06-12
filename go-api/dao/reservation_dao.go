package dao

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	HotelID int `gorm:"column:hotel_id"`
	//Hotel      Hotel
	CheckIn    string `gorm:"column:checkout"`
	CheckOut   string `gorm:"column:checkin"`
	ClientName string `gorm:"column:cliente_nombre"`
}
