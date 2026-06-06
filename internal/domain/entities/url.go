package entities

import (
	"fmt"
	"math/big"
	"net/url"

	"github.com/clevanilson/cs-url-shortner/pkg/c_errors"
)

type URL struct {
	original string
	shorten  string
}

func NewURL(original string, id uint64) (*URL, error) {
	if a, err := url.ParseRequestURI(original); err != nil {
		fmt.Printf("%v", a)
		return nil, c_errors.NewDomainError("Invalid URL")
	}
	return &URL{
		original: original,
		shorten:  big.NewInt(int64(id)).Text(62),
	}, nil
}

func (u *URL) Original() string {
	return u.original
}

func (u *URL) Shorten() string {
	return u.shorten
}
