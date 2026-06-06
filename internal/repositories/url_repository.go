package repositories

import (
	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/pkg/c_errors"
)

type IURLRepository interface {
	Save(url *entities.URL) error
	GetByShortURL(shorten string) (*entities.URL, error)
}

type UrlMemoryRepository struct {
	data map[string]*entities.URL
}

func NewURLMemoryRepository() *UrlMemoryRepository {
	return &UrlMemoryRepository{
		data: make(map[string]*entities.URL),
	}
}

func (r *UrlMemoryRepository) Save(url *entities.URL) error {
	r.data[url.Shorten()] = url
	return nil
}

func (r *UrlMemoryRepository) GetByShortURL(shorten string) (*entities.URL, error) {
	url, ok := r.data[shorten]
	if !ok {
		return nil, c_errors.NewNotFoundError(shorten)
	}
	return url, nil
}
