package dao

type HotelPhoto struct {
	ID       int    `gorm:"primaryKey"`
	HotelID  int    `gorm:"column:hotel_id"`
	Filename string `gorm:"column:filename"`
}

func (HotelPhoto) TableName() string {
	return "photos"
}
