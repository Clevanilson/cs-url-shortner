package repositories

import "github.com/clevanilson/cs-url-shortner/internal/domain/entities"

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
	url := r.data[shorten]
	return url, nil
}
