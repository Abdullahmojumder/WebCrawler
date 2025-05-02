package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		expectError bool
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Other</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			expectError: false,
		},
		{
			name:     "multiple relative URLs and empty href",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="/about">About</a>
		<a href="/contact/us">Contact</a>
		<a href="">Empty</a>
		<a>Not a link</a>
	</body>
</html>
`,
			expected:    []string{"https://example.com/about", "https://example.com/contact/us"},
			expectError: false,
		},
		{
			name:     "invalid base URL",
			inputURL: "://invalid",
			inputBody: `
<html>
	<body>
		<a href="/path">Path</a>
	</body>
</html>
`,
			expected:    nil,
			expectError: true,
		},
		{
			name:     "mixed URLs with fragments and queries",
			inputURL: "http://boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path?key=value">Query</a>
		<a href="/path#section">Fragment</a>
		<a href="https://test.com/path">Absolute</a>
	</body>
</html>
`,
			expected:    []string{"http://boot.dev/path?key=value", "http://boot.dev/path#section", "https://test.com/path"},
			expectError: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if (err != nil) != tc.expectError {
				t.Errorf("Test %v - '%s' FAIL: unexpected error status: got error %v, expectError %v", i, tc.name, err, tc.expectError)
				return
			}
			// Sort slices for consistent comparison
			if actual != nil {
				sort.Strings(actual)
			}
			if tc.expected != nil {
				sort.Strings(tc.expected)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
