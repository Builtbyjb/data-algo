package main

import "testing"

func TestFound(t *testing.T) {
	haystack := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var needle int16 = 5

	result := binary_search(haystack, needle)
	if result != true {
		t.Errorf("Expected true")
	}
}

func TestNotFound(t *testing.T) {
	haystack := []int16{1, 2, 4, 6, 7, 8, 9, 11, 14, 18}
	var needle int16 = 5

	result := binary_search(haystack, needle)
	if result != false {
		t.Errorf("Expected false")
	}
}
