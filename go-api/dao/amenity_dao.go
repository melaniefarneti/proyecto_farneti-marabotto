package dao

type Amenity struct {
	ID      int    `gorm:"type:int;primaryKey"`
	Nombre  string `gorm:"type:varchar(50);column:nombre"`
	HotelID int64  `gorm:"column:hotel_id"`
}
