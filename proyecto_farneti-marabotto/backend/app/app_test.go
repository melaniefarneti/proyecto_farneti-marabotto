package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPort(t *testing.T) {  //verifica si el valor de la constante port es ":8080".
	assert.Equal(t, ":8080", port)
}