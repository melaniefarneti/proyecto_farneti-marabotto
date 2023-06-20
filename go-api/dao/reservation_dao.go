package dao

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	HotelID int `gorm:"column:hotel_id"`
	//Hotel      Hotel
	CheckIn    string `gorm:"column:checkin"`
	CheckOut   string `gorm:"column:checkout"`
	ClientName string `gorm:"column:cliente_nombre"`
	UserID     int    `gorm:"column:user_id"`
}
