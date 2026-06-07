package repositories

import (
	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/pkg/c_errors"
	"github.com/clevanilson/cs-url-shortner/pkg/database"
)

type IURLRepository interface {
	Save(url *entities.URL) error
	GetByShortURL(shorten string) (*entities.URL, error)
	GetNextId() (int64, error)
}

type UrlMemoryRepository struct {
	kvConnection database.KeyValueConnection
	data         map[string]*entities.URL
}

func NewURLMemoryRepository(kvConnection database.KeyValueConnection) *UrlMemoryRepository {
	return &UrlMemoryRepository{
		data:         make(map[string]*entities.URL),
		kvConnection: kvConnection,
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

func (r *UrlMemoryRepository) GetNextId() (int64, error) {
	id, err := r.kvConnection.Increment("global:id")
	if err != nil {
		return 0, err
	}
	return id, nil
}
