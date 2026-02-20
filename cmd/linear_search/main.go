package main

import (
	"fmt"
	"slices"
)

/* Linear search searching each element in the array one by one
* Time complexity O(N)
 */
func linear_search(haystack []int16, needle int16) bool {
	// for _, i := range haystack {
	// 	if needle == i {
	// 		return true
	// 	}
	// }
	return slices.Contains(haystack, needle)
}

func main() {
	haystack := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := linear_search(haystack, 5)
	fmt.Println(result)
}
