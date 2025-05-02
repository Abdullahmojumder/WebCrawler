package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// getHTML fetches the HTML content of the given URL and returns it as a string.
func getHTML(rawURL string) (string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Check for error-level status codes (400+)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	// Check the Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(strings.ToLower(contentType), "text/html") {
		return "", fmt.Errorf("content type is not text/html: %s", contentType)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
