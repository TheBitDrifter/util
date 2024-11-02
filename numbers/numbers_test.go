package numbers_util

import (
	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "With duplicates",
			input:    []int{1, 2, 2, 3, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "All duplicates",
			input:    []int{5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "Mixed values",
			input:    []int{2, 3, 1, 3, 2, 4, 1},
			expected: []int{2, 3, 1, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UniqueInts(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestDescendingInts(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Already sorted descending",
			input:    []int{3, 2, 1},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Sorted ascending",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "With duplicates",
			input:    []int{4, 2, 3, 1, 2},
			expected: []int{4, 3, 2, 2, 1},
		},
		{
			name:     "All duplicates",
			input:    []int{5, 5, 5},
			expected: []int{5, 5, 5},
		},
		{
			name:     "Mixed values",
			input:    []int{2, 3, 1, 4},
			expected: []int{4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DescendingInts(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
