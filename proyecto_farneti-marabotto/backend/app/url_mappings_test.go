package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
	assert.Equal(t, "/items/:itemID", pathGetItem) //verifica si el valor de pathGetItem
	                                               // es igual a "/items/:itemID".
}