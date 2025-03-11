package adapters

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"wichipu.com/scraper/internal/domain"
)

func TestWebScraper_Scrape(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
			<body>
				<a href="/page1">Link 1</a>
				<a href="/page2">Link 2</a>
				<a href="https://external.com">External Link</a>
			</body>
			</html>
		`))
	}))
	defer ts.Close()

	// Create the test site
	site, err := domain.NewSite(ts.URL)
	if err != nil {
		t.Fatalf("Error creating site: %v", err)
	}

	// Create the scraper and configure it to allow the test server domain
	scraper := NewWebScraper()

	// Execute scraping
	work, err := scraper.Scrape(context.Background(), site)
	if err != nil {
		t.Fatalf("Error during scraping: %v", err)
	}

	// Verify results
	assert.NotNil(t, work)
	assert.NotZero(t, work.StartedAt)
	assert.NotNil(t, work.FinishedAt)
	assert.True(t, work.FinishedAt.After(work.StartedAt))
}
