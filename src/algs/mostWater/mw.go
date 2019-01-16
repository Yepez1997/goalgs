package main

import (
	"fmt"
	"math"
)

// find two lines x1, x2 in array 1 .. n where  1 <= x1 <= x2 <= n
// such that x1 and x2 contains the most water
// each index represents a certain arr

func container(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	// max area
	var max float64
	i := 0
	j := len(arr) - 1

	for i <= j {
		if arr[i] < arr[j] {
			max = math.Max(max, float64(arr[i]*(j-i)))
			i++
		} else {
			// value of j is greater
			max = math.Max(max, float64(arr[j]*(j-i)))
			j--
		}
	}
	return int(max)
}

func main() {
	// [1,8,6,2,5,4,8,3,7] -> 49
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := container(input)
	fmt.Println(result)

}

/*
	heap, stack
	maintain the current largest distance
	waterheld = 0
	j = 1
	for i ... n - 1
		if arr[i] < arr[j]
			area = arr[i] *  arr[j]
			if area > waterhelp:
				waterheld = area

		j++


*/
