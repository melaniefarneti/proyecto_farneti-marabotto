package main

import (
	"go-api/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Configura la conexión a la base de datos
	dsn := "root:mel1@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Obtiene la instancia de *sql.DB desde el objeto *gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()
	// Prueba la conexión
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app.StartApp()
}
