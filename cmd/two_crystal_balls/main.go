package main

import (
	"fmt"
	"math"
)

// Time complexity O(sqrt(N))
func twoCrystalBalls(breaks []bool) int8 {
	len := len(breaks)
	jmpAmount := math.Floor(math.Sqrt(float64(len)))

	i := int(jmpAmount)

	//  Get when first ball breaks
	for i < len {
		if breaks[i] {
			break
		}

		i += int(jmpAmount)
	}

	// Jump back by one sqrt
	i -= int(jmpAmount)

	j := 0
	// Get exact break point
	for j < int(jmpAmount) && i < len-1 {
		i++
		j++

		if breaks[i] {
			return int8(i)
		}
	}

	return -1
}

func main() {
	breaks := []bool{false, false, false, false, false, false, true, true, true}
	result := twoCrystalBalls(breaks)
	fmt.Println(result) // Ans 6
}
