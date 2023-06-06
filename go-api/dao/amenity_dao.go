package dao

import "gorm.io/gorm"

type Amenity struct {
	gorm.Model
	Name    string `gorm:"type:longtext;column:nombre"`
	HotelID int    `gorm:"type:int not null"`
	Hotel   Hotel
}
