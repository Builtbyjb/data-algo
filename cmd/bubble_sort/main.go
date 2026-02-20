package main

import "fmt"

// Time complexity O(N^2)
func bubbleSort(arr *[]int16) {
	len := len(*arr)

	for i := range len {
		for j := range len - 1 - i {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
				// Quick swap
				// (*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	arr := []int16{3, 6, 1, 2, 7, 9, 4, 5, 6, 8}
	bubbleSort(&arr)

	fmt.Println(arr)
}
