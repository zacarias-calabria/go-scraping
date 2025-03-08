package domain

import "errors"

type Link struct {
	URL        string            `json:"url"`
	Content    string            `json:"content"`
	Attributes map[string]string `json:"attributes"`
}

func NewLink(URL string, content string, attributes map[string]string) (*Link, error) {
	if URL == "" {
		return nil, errors.New("error.link.url_is_empty")
	}
	if content == "" {
		return nil, errors.New("error.link.content_is_empty")
	}
	return &Link{URL: URL, Content: content, Attributes: attributes}, nil
}
