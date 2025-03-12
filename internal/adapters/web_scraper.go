package adapters

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"wichipu.com/scraper/internal/domain"
	"wichipu.com/scraper/internal/ports"
)

type WebScraper struct {
	collector *colly.Collector
}

// userAgents es una lista de User-Agents comunes para rotar
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Safari/605.1.15",
}

func NewWebScraper() ports.Scrapper {
	// Configuración básica del collector
	c := colly.NewCollector(
		colly.MaxDepth(5), // Default maximum depth
		colly.Async(true), // Habilitar scraping asíncrono
	)

	// Configurar límites de velocidad
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second,
		Parallelism: 2,
	})

	// Basic collector configuration
	c.SetRequestTimeout(30 * time.Second)
	c.AllowURLRevisit = false

	// Configurar cabeceras por defecto
	c.OnRequest(func(r *colly.Request) {
		// Rotar User-Agent aleatoriamente
		r.Headers.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		//r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

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

	// Configurar proxy si está disponible en las variables de entorno
	if proxyURL := os.Getenv("HTTP_PROXY"); proxyURL != "" {
		w.collector.SetProxy(proxyURL)
	}

	// Handle found links
	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// TODO: Implement logic to save found links
		fmt.Printf("Found link: %s\n", link)
	})

	// Manejar respuestas
	w.collector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK {
			fmt.Printf("Got status code: %d for %s\n", r.StatusCode, r.Request.URL)
		}
	})

	// Handle errors
	w.collector.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == http.StatusForbidden {
			fmt.Printf("Access forbidden (403) for %s - might need to adjust request headers or respect robots.txt\n", r.Request.URL)
		} else {
			fmt.Printf("Error scraping %s: %v\n", r.Request.URL, err)
		}
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
