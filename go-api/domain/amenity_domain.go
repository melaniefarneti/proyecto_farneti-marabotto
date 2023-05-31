package domain

type Amenity struct {
	ID      int    `gorm:"type:int;primaryKey"`
	Name    string `gorm:"type:longtext;column:nombre"`
	HotelID int    `gorm:"type:int not null;column:hotel_id"`
}
