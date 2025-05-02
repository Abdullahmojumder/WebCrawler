package main

import (
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

// getURLsFromHTML extracts all URLs from <a> tags in the HTML body, converting relative URLs to absolute using rawBaseURL.
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	// Parse the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	// Create a reader from the HTML body
	reader := strings.NewReader(htmlBody)

	// Parse the HTML into a node tree
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	// Collect URLs recursively
	var urls []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && attr.Val != "" {
					// Parse the href value
					parsedURL, err := url.Parse(attr.Val)
					if err != nil {
						continue // Skip invalid URLs
					}
					// Convert relative URLs to absolute
					if !parsedURL.IsAbs() {
						parsedURL = baseURL.ResolveReference(parsedURL)
					}
					urls = append(urls, parsedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return urls, nil
}
