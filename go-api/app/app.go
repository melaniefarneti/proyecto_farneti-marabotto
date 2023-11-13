package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	// Configurar el middleware CORS con la configuración adecuada
	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	mapRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Errorf("error running server: %w", err))
	}
}
