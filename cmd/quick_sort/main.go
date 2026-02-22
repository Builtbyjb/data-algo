package main

import "fmt"

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {

	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{7, 8, 9, 2, 1, 8, 6, 9, 11, 2, 14}
	quickSort(arr)
	fmt.Println(arr)
}
