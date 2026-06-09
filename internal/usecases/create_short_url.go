package usecases

import (
	"fmt"

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
	url, err := u.repostory.GetByOriginalURL(input.LongUrl)
	if err != nil {
		return nil, err
	}
	if url != nil {
		return formatOutput(url), nil
	}
	id, err := u.repostory.GetNextId()
	if err != nil {
		return nil, err
	}
	url, err = entities.NewURL(input.LongUrl, id)
	if err != nil {
		return nil, err
	}
	if err := u.repostory.Save(url); err != nil {
		return nil, err
	}
	return formatOutput(url), nil
}

func formatOutput(url *entities.URL) *CreateShortUrlOutput {
	return &CreateShortUrlOutput{
		ShortUrl: fmt.Sprintf("http://localhost:3000/%v", url.Shorten()),
		Code:     url.Shorten(),
	}
}

type CreateShortUrlInput struct {
	LongUrl string `json:"long_url"`
}

type CreateShortUrlOutput struct {
	ShortUrl string `json:"short_url"`
	Code     string `json:"code"`
}
