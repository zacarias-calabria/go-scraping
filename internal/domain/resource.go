package domain

import "errors"

type Resource struct {
	URL      string `json:"url"`
	MIMEType string `json:"mime_type"`
	Referer  string `json:"referer"`
	Content  string `json:"content"`
}

func NewResource(URL string, MIMEType string, referer string, content string) (*Resource, error) {
	if URL == "" {
		return nil, errors.New("error.resource.url_is_empty")
	}
	if MIMEType == "" {
		return nil, errors.New("error.resource.mime_type_is_empty")
	}
	if referer == "" {
		return nil, errors.New("error.resource.referer_is_empty")
	}
	if content == "" {
		return nil, errors.New("error.resource.content_is_empty")
	}
	return &Resource{URL: URL, MIMEType: MIMEType, Referer: referer, Content: content}, nil
}
