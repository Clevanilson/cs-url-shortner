package usecases_test

import (
	"testing"

	"github.com/clevanilson/cs-url-shortner/internal/repositories"
	"github.com/clevanilson/cs-url-shortner/internal/usecases"
	"github.com/clevanilson/cs-url-shortner/pkg/assert"
)

func TestCreateShortUrl(t *testing.T) {
	repository := repositories.NewURLMemoryRepository()
	sut := usecases.NewCreateShortUrl(repository)

	t.Run("With valid data", func(t *testing.T) {
		input := usecases.CreateShortUrlInput{LongUrl: "https://google.com"}
		output, err := sut.Execute(input)
		assert.Equal(t, err, nil)
		savedUrl, err := repository.GetByShortURL(output.ShortUrl)
		assert.Equal(t, err, nil)
		assert.Equal(t, savedUrl.Original(), input.LongUrl)
	})

	t.Run("With invalid data", func(t *testing.T) {
		input := usecases.CreateShortUrlInput{LongUrl: "invalid"}
		output, err := sut.Execute(input)
		assert.Equal(t, output, nil)
		assert.Equal(t, err.Error(), "[DomainError]: Invalid URL")
	})
}
