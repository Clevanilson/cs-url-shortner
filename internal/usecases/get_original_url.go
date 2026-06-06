package usecases

import (
	"github.com/clevanilson/cs-url-shortner/internal/repositories"
	"github.com/clevanilson/cs-url-shortner/pkg/c_errors"
)

type GetOriginalUrl struct {
	repository repositories.IURLRepository
}

func NewGetOriginaUrl(repository repositories.IURLRepository) *GetOriginalUrl {
	return &GetOriginalUrl{repository}
}

func (u *GetOriginalUrl) Execute(input GetOriginalUrlInput) (*GetOriginalUrlOutput, error) {
	url, err := u.repository.GetByShortURL(input.ShortUrl)
	if err != nil {
		return nil, c_errors.NewNotFoundError(input.ShortUrl)
	}
	return &GetOriginalUrlOutput{OriginalUrl: url.Original()}, nil
}

type GetOriginalUrlInput struct {
	ShortUrl string
}

type GetOriginalUrlOutput struct {
	OriginalUrl string
}
