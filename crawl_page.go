package main

import (
	"fmt"
	"net/url"
	"sync"
)

// config holds shared state for the crawler.
type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

// addPageVisit adds a normalized URL to the pages map and returns true if it's the first visit.
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if count, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL] = count + 1
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

// crawlPage recursively crawls a website starting from rawCurrentURL, using goroutines.
func (cfg *config) crawlPage(rawCurrentURL string) {
	// Check if maxPages limit is reached
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	// Parse the current URL
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing current URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// Check if current URL is on the same domain as base URL
	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	// Normalize the current URL
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	// Add page visit and check if it's the first time
	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	// Fetch the HTML
	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error fetching HTML from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Extract URLs from the HTML
	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Recursively crawl each URL in a new goroutine
	for _, nextURL := range urls {
		cfg.wg.Add(1)
		go func(url string) {
			cfg.concurrencyControl <- struct{}{} // Acquire concurrency slot
			defer cfg.wg.Done()
			defer func() { <-cfg.concurrencyControl }() // Release concurrency slot
			cfg.crawlPage(url)
		}(nextURL)
	}
}
