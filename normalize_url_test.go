package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		expectError   bool
	}{
		{
			name:        "remove scheme",
			inputURL:    "https://blog.boot.dev/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "remove scheme http",
			inputURL:    "http://blog.boot.dev/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "remove trailing slash",
			inputURL:    "https://blog.boot.dev/path/",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "handle no scheme",
			inputURL:    "blog.boot.dev/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "lowercase hostname",
			inputURL:    "https://BLOG.BOOT.DEV/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "remove default port",
			inputURL:    "https://blog.boot.dev:443/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "remove default http port",
			inputURL:    "http://blog.boot.dev:80/path",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "preserve non-default port",
			inputURL:    "https://blog.boot.dev:8080/path",
			expected:    "blog.boot.dev:8080/path",
			expectError: false,
		},
		{
			name:        "remove fragment",
			inputURL:    "https://blog.boot.dev/path#section",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "remove query params",
			inputURL:    "https://blog.boot.dev/path?key=value",
			expected:    "blog.boot.dev/path",
			expectError: false,
		},
		{
			name:        "handle empty path",
			inputURL:    "https://blog.boot.dev",
			expected:    "blog.boot.dev",
			expectError: false,
		},
		{
			name:        "handle root path",
			inputURL:    "https://blog.boot.dev/",
			expected:    "blog.boot.dev",
			expectError: false,
		},
		{
			name:        "invalid URL",
			inputURL:    "://invalid",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty URL",
			inputURL:    "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "path with multiple slashes",
			inputURL:    "https://blog.boot.dev/path//to//resource",
			expected:    "blog.boot.dev/path/to/resource",
			expectError: false,
		},
		{
			name:        "handle subdomains",
			inputURL:    "https://sub.blog.boot.dev/path",
			expected:    "sub.blog.boot.dev/path",
			expectError: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if (err != nil) != tc.expectError {
				t.Errorf("Test %v - '%s' FAIL: unexpected error status: got error %v, expectError %v", i, tc.name, err, tc.expectError)
				return
			}
			if err == nil && actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
