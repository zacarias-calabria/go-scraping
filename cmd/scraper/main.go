package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"wichipu.com/scraper/internal/adapters"
	"wichipu.com/scraper/internal/domain"
)

func main() {
	// Define command line flags
	url := flag.String("url", "", "URL to scrape")
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout for the scraping process")
	proxy := flag.String("proxy", "", "Proxy URL (e.g., http://proxy:port)")
	flag.Parse()

	if *url == "" {
		fmt.Println("Error: URL is required")
		flag.Usage()
		os.Exit(1)
	}

	// Configure proxy if provided
	if *proxy != "" {
		os.Setenv("HTTP_PROXY", *proxy)
	}

	// Create site
	site, err := domain.NewSite(*url)
	if err != nil {
		fmt.Printf("Error creating site: %v\n", err)
		os.Exit(1)
	}

	// Create scraper
	scraper := adapters.NewWebScraper()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	fmt.Printf("Starting scraping of %s\n", *url)
	if *proxy != "" {
		fmt.Printf("Using proxy: %s\n", *proxy)
	}

	// Execute scraping
	work, err := scraper.Scrape(ctx, site)
	if err != nil {
		fmt.Printf("Error during scraping: %v\n", err)
		os.Exit(1)
	}

	// Print results
	fmt.Printf("\nScraping completed successfully\n")
	fmt.Printf("Started at: %v\n", work.StartedAt)
	fmt.Printf("Finished at: %v\n", *work.FinishedAt)
	fmt.Printf("Duration: %v\n", work.FinishedAt.Sub(work.StartedAt))
}
