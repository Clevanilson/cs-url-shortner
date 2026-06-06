package usecases_test

import (
	"testing"

	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/internal/repositories"
	"github.com/clevanilson/cs-url-shortner/internal/usecases"
	"github.com/clevanilson/cs-url-shortner/pkg/assert"
)

func TestGetOriginalUrl(t *testing.T) {
	repository := repositories.NewURLMemoryRepository()
	sut := usecases.NewGetOriginaUrl(repository)

	t.Run("With existing url", func(t *testing.T) {
		url, _ := entities.NewURL("https://google.com", 100000)
		repository.Save(url)
		output, err := sut.Execute(usecases.GetOriginalUrlInput{ShortUrl: url.Shorten()})
		assert.Equal(t, err, nil)
		assert.Equal(t, output.OriginalUrl, url.Original())
	})

	t.Run("With non-existent url", func(t *testing.T) {
		url, _ := entities.NewURL("https://youtube.com", 123404)
		output, err := sut.Execute(usecases.GetOriginalUrlInput{ShortUrl: url.Shorten()})
		assert.Equal(t, output, nil)
		assert.Equal(t, err.Error(), "[NotFound]: "+url.Shorten())
	})
}
