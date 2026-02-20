package main

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int16{3, 6, 1, 2, 7, 9, 4, 5, 6, 8}
	sorted := []int16{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}

	bubbleSort(&arr)
	if !slices.Equal(arr, sorted) {
		t.Errorf("Expected arr == sorted")
	}
}

func TestSortTwo(t *testing.T) {
	arr := []int16{3, 6, 1, 2, 7, 9, 4, 5, 6, 8, 11, 19, 2, 8, 14}
	sorted := []int16{1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 11, 14, 19}

	bubbleSort(&arr)
	if !slices.Equal(arr, sorted) {
		t.Errorf("Expected arr == sorted")
	}
}
