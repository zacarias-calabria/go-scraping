package domain

import "errors"

type Site struct {
	URL string `json:"url"`
}

func NewSite(url string) (*Site, error) {
	if url == "" {
		return nil, errors.New("error.site.url_is_empty")
	}
	return &Site{URL: url}, nil
}
