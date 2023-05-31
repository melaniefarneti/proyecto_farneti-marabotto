package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLs(t *testing.T) {
	assert.Equal(t, "/hotels/%d", hotelEndpoint)
	assert.Equal(t, "https://example.com/api%s", baseURL)
}
