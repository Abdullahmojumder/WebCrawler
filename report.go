package main

import (
	"fmt"
	"sort"
)

// pageEntry holds a URL and its link count for sorting.
type pageEntry struct {
	url   string
	count int
}

// sortPages sorts the pages map into a slice of pageEntry, ordered by count (descending) and then URL (alphabetically).
func sortPages(pages map[string]int) []pageEntry {
	entries := make([]pageEntry, 0, len(pages))
	for url, count := range pages {
		entries = append(entries, pageEntry{url: url, count: count})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count == entries[j].count {
			return entries[i].url < entries[j].url
		}
		return entries[i].count > entries[j].count
	})
	return entries
}

// printReport prints a formatted report of the crawled pages.
func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")

	entries := sortPages(pages)
	for _, entry := range entries {
		fmt.Printf("Found %d internal links to %s\n", entry.count, entry.url)
	}
}
