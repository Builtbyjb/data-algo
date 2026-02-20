package main

import (
	"fmt"
)

/*
	NOTE: Primary conditions for binary search

Time complexity O(log N)
- Value equals target
- Value is larger than target
- Value is smaller than target
*/
func binary_search(haystack []int16, needle int16) bool {
	low := 0
	high := len(haystack)

	for low < high {
		mid := low + ((high - low) / 2)

		value := haystack[mid]

		if value == needle {
			return true
		} else if value > needle {
			high = mid
		} else if value < needle {
			low = mid + 1
		}
	}

	return false
}

func main() {
	haystack := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var needle int16 = 5
	result := binary_search(haystack, needle)

	fmt.Println(result)
}
