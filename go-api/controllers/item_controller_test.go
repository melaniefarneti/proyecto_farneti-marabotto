package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*verifica si el valor de paramItemID es igual a "itemID" */
func TestParams(t *testing.T) {
	assert.Equal(t, "itemID", paramItemID)
}