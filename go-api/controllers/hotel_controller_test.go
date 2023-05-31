package controllers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	// Crea un contexto de prueba
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(nil)

	// Llama a la función GetHotels del controlador
	GetHotels(ctx)

	// Verifica el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
