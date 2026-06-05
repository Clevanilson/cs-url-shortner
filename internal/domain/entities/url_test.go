package entities_test

import (
	"testing"

	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/pkg/assert"
)

func TestURL(t *testing.T) {
	t.Run("Encoding URL", func(t *testing.T) {
		original := "https://google.com"
		url, err := entities.NewURL(original, 11157)
		assert.Equal(t, err, nil)
		assert.Equal(t, url.Original(), original)
		assert.Equal(t, url.Shorten(), "2TX")
	})

	t.Run("With invalid URL", func(t *testing.T) {
		url, err := entities.NewURL("invalid", 13344)
		assert.Equal(t, url, nil)
		assert.Equal(t, err.Error(), "Invalid URL")
	})
}
