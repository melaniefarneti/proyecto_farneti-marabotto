package dao

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	ID          int    `gorm:"type:int;primaryKey"`
	Name        string `gorm:"type:varchar(255);column:nombre"`
	Photo       string `gorm:"type:varchar(255);column:foto"`
	Description string `gorm:"type:varchar(1000);column:descripcion"`
	Location    string `gorm:"type:varchar(50);column:ubicacion"`
	Rooms       string `gorm:"type:varchar(255);column:cuartos"`
}
