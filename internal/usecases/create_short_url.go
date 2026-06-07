package usecases

import (
	"github.com/clevanilson/cs-url-shortner/internal/domain/entities"
	"github.com/clevanilson/cs-url-shortner/internal/repositories"
)

type CreateShortUrl struct {
	repostory repositories.IURLRepository
}

func NewCreateShortUrl(repository repositories.IURLRepository) *CreateShortUrl {
	return &CreateShortUrl{repository}
}

func (u *CreateShortUrl) Execute(input CreateShortUrlInput) (*CreateShortUrlOutput, error) {
	id, err := u.repostory.GetNextId()
	if err != nil {
		return nil, err
	}
	url, err := entities.NewURL(input.LongUrl, id)
	if err != nil {
		return nil, err
	}
	if err := u.repostory.Save(url); err != nil {
		return nil, err
	}
	return &CreateShortUrlOutput{ShortUrl: url.Shorten()}, nil
}

type CreateShortUrlInput struct {
	LongUrl string `json:"long_url"`
}

type CreateShortUrlOutput struct {
	ShortUrl string `json:"short_url"`
}
