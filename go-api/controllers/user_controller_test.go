package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserParams(t *testing.T) {
	assert.Equal(t, "userID", paramUserID)
}
