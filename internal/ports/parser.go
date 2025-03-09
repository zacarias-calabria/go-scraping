package ports

import (
	"context"

	"wichipu.com/scraper/internal/domain"
)

type Scrapper interface {
	Scrape(ctx context.Context, url string) (*domain.Work, error)
}
