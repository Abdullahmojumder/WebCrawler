package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:] // Exclude the program name
	if len(args) != 3 {
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	baseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a positive integer")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s (maxConcurrency: %d, maxPages: %d)\n", baseURL, maxConcurrency, maxPages)

	// Parse base URL
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		os.Exit(1)
	}

	// Initialize config
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	// Start crawling
	cfg.wg.Add(1)
	go func() {
		cfg.concurrencyControl <- struct{}{}
		defer cfg.wg.Done()
		defer func() { <-cfg.concurrencyControl }()
		cfg.crawlPage(baseURL)
	}()

	// Wait for all goroutines to complete
	cfg.wg.Wait()

	// Print report
	printReport(cfg.pages, baseURL)
}
