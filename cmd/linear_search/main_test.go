package main

import "testing"

func TestLinearSearchFound(t *testing.T) {
	haystack := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := linear_search(haystack, 6)

	if result != true {
		t.Errorf("Expected true")
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	haystack := []int16{1, 11, 3, 4, 5, 6, 7, 8, 9}
	result := linear_search(haystack, 2)

	if result != false {
		t.Errorf("Expected false")
	}
}
