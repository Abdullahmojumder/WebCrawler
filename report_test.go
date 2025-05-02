package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []pageEntry
	}{
		{
			name: "sort by count descending then URL alphabetically",
			input: map[string]int{
				"example.com/page1": 3,
				"example.com/page2": 2,
				"example.com/page3": 3,
				"example.com/page4": 1,
			},
			expected: []pageEntry{
				{url: "example.com/page1", count: 3},
				{url: "example.com/page3", count: 3},
				{url: "example.com/page2", count: 2},
				{url: "example.com/page4", count: 1},
			},
		},
		{
			name: "same counts sorted alphabetically",
			input: map[string]int{
				"example.com/zebra":  2,
				"example.com/apple":  2,
				"example.com/banana": 2,
			},
			expected: []pageEntry{
				{url: "example.com/apple", count: 2},
				{url: "example.com/banana", count: 2},
				{url: "example.com/zebra", count: 2},
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: []pageEntry{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
