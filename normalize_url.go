// normalize_url.go
package main

import (
	"net/url"
	"errors"
	"strings"
)

// normalizeURL takes a URL string and returns a normalized version of it.
func normalizeURL(rawURL string) (string, error) {
	// Handle empty input
	if rawURL == "" {
		return "", &url.Error{Op: "parse", URL: rawURL, Err: errors.New("empty URL")}
	}

	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// If no scheme is provided, Parse may not set Host correctly; try adding a scheme
	if parsedURL.Host == "" && !strings.Contains(rawURL, "://") {
		parsedURL, err = url.Parse("http://" + rawURL)
		if err != nil {
			return "", err
		}
	}

	// Get the hostname (lowercase) and port
	host := strings.ToLower(parsedURL.Hostname())
	port := parsedURL.Port()

	// Remove default ports (80 for HTTP, 443 for HTTPS)
	if (parsedURL.Scheme == "http" && port == "80") || (parsedURL.Scheme == "https" && port == "443") || port == "" {
		port = ""
	}

	// Build the host part (hostname + port if non-empty)
	hostPart := host
	if port != "" {
		hostPart = host + ":" + port
	}

	// Normalize the path: remove trailing slash and collapse multiple slashes
	path := parsedURL.Path
	if path == "" {
		path = "/"
	}
	// Remove trailing slash
	path = strings.TrimSuffix(path, "/")
	// Collapse multiple slashes
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}

	// If path is just "/", omit it in the output
	if path == "/" {
		path = ""
	}

	// Combine host and path
	normalized := hostPart
	if path != "" {
		normalized += path
	}

	return normalized, nil
}
