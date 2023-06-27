package app

import (
	"html/template"

	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

// función auxiliar para cargar y analizar el archivo de plantilla
func loadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

func StartApp() {
	router := gin.Default()
	router.Use(corsMiddleware()) // Agregar el middleware CORS

	mapRoutes(router)

	err := router.Run(port)
	if err != nil {
		panic(fmt.Errorf("error running app: %w", err))
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	}
}
