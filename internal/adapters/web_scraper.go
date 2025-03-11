package adapters

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly/v2"
	"wichipu.com/scraper/internal/domain"
	"wichipu.com/scraper/internal/ports"
)

type WebScraper struct {
	collector *colly.Collector
}

func NewWebScraper() ports.Scrapper {
	c := colly.NewCollector(
		colly.MaxDepth(5), // Default maximum depth
	)

	// Basic collector configuration
	c.SetRequestTimeout(30 * time.Second)
	c.AllowURLRevisit = false

	return &WebScraper{
		collector: c,
	}
}

func (w *WebScraper) Scrape(ctx context.Context, site *domain.Site) (*domain.Work, error) {
	work := domain.NewWork(time.Now())

	// Configure collector based on the site
	siteURL, err := url.Parse(site.URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing site URL: %w", err)
	}

	// Allow site domain and localhost for testing
	w.collector.AllowedDomains = []string{siteURL.Host, "127.0.0.1", "localhost"}

	// Handle found links
	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// TODO: Implement logic to save found links
		fmt.Printf("Found link: %s\n", link)
	})

	// Handle errors
	w.collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error scraping %s: %v\n", r.Request.URL, err)
	})

	// Start scraping
	err = w.collector.Visit(site.URL)
	if err != nil {
		return nil, fmt.Errorf("error starting scraping: %w", err)
	}

	// Wait for scraping to finish
	w.collector.Wait()

	// Mark work as finished
	err = work.Finish(time.Now())
	if err != nil {
		return nil, fmt.Errorf("error finishing work: %w", err)
	}

	return work, nil
}
