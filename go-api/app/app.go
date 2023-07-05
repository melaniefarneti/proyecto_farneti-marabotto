package app

import (
	"fmt"
	"html/template"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func loadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

func StartApp() {
	router := gin.Default()

	// Configura el middleware CORS con la configuración adecuada
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"*"} // Permitir todas las cabeceras
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// Ruta de prueba
	router.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Login successful"})
	})

	err := router.Run(port)
	if err != nil {
		panic(fmt.Errorf("error running app: %w", err))
	}
}
