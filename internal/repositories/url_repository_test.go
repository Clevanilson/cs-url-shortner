package repositories_test

import (
	"testing"

	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/internal/repositories"
	"github.com/clevanilson/cs-url-shortner/pkg/assert"
)

func TestURLRepository(t *testing.T) {
	sut := repositories.NewURLMemoryRepository()
	url, _ := entities.NewURL("http://google.com", 100000)

	t.Run("Saving url", func(t *testing.T) {
		err := sut.Save(url)
		assert.Equal(t, err, nil)
	})

	t.Run("Retriving url by shorten url", func(t *testing.T) {
		sut.Save(url)
		savedUrl, err := sut.GetByShortURL(url.Shorten())
		assert.Equal(t, err, nil)
		assert.Equal(t, savedUrl, url)
	})

	t.Run("Trying to get a non-existent url", func(t *testing.T) {
		savedUrl, err := sut.GetByShortURL("TX901")
		assert.Equal(t, savedUrl, nil)
		assert.Equal(t, err.Error(), "[NotFound]: TX901")
	})

}
