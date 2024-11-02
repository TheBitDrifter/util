package iter_util

import (
	"iter"
	"testing"
)

func TestCollect(t *testing.T) {
	// Create a test sequence using a generator function
	testSeq := iter.Seq[int](func(yield func(int) bool) {
		for _, v := range []int{1, 2, 3, 4, 5} {
			if !yield(v) {
				return
			}
		}
	})

	// Call the Collect function
	result := Collect(testSeq)

	// Expected result
	expected := []int{1, 2, 3, 4, 5}

	// Check if the result matches the expected output
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, but got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("At index %d, expected %d, but got %d", i, expected[i], v)
		}
	}
}

func TestCollect2(t *testing.T) {
	// Create a test sequence using a generator function
	testSeq2 := iter.Seq2[int, string](func(yield func(int, string) bool) {
		for _, pair := range []struct {
			id   int
			name string
		}{
			{1, "one"},
			{2, "two"},
			{3, "three"},
		} {
			if !yield(pair.id, pair.name) {
				return
			}
		}
	})

	// Call the Collect2 function
	result := Collect2(testSeq2)

	// Expected result
	expected := []string{"one", "two", "three"}

	// Check if the result matches the expected output
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, but got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("At index %d, expected %s, but got %s", i, expected[i], v)
		}
	}
}
