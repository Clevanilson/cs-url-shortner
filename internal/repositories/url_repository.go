package repositories

import (
	"errors"

	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/pkg/database"
	"github.com/gocql/gocql"
)

type IURLRepository interface {
	Save(url *entities.URL) error
	GetByShortURL(shorten string) (*entities.URL, error)
	GetByOriginalURL(shorten string) (*entities.URL, error)
	GetNextId() (int64, error)
}

type UrlDatabaseRepository struct {
	kvConnection        database.RedisCoonection
	cassandraConnection database.CassandraConnection
}

func NewUrlDatabaseRepository(
	kvConnection database.RedisCoonection,
	cassandraConnection database.CassandraConnection,
) *UrlDatabaseRepository {
	return &UrlDatabaseRepository{
		kvConnection:        kvConnection,
		cassandraConnection: cassandraConnection,
	}
}

func (r *UrlDatabaseRepository) Save(url *entities.URL) error {
	err := r.cassandraConnection.Exec(`
		INSERT INTO url_shortner.urls(short, original)
		VALUES (?, ?);
		`,
		url.Shorten(),
		url.Original(),
	)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *UrlDatabaseRepository) GetByShortURL(shorten string) (*entities.URL, error) {
	return r.getUrl(
		`SELECT short, original FROM url_shortner.urls WHERE short = ?;`,
		shorten,
	)
}

func (r *UrlDatabaseRepository) GetByOriginalURL(long string) (*entities.URL, error) {
	return r.getUrl(
		`SELECT short, original
		FROM url_shortner.urls WHERE original = ? ALLOW FILTERING;`,
		long,
	)
}

func (r *UrlDatabaseRepository) GetNextId() (int64, error) {
	id, err := r.kvConnection.Increment("global:id")
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UrlDatabaseRepository) getUrl(query string, arg string) (*entities.URL, error) {
	_query := r.cassandraConnection.Query(query, arg)
	var short string
	var original string
	if err := _query.Scan(&short, &original); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	url, err := entities.Create(original, short)
	if err != nil {
		return nil, err
	}
	return url, nil
}
