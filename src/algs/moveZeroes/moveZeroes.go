/*Given an array nums, write a function to move all 0's to the end of it
while maintaining the relative order of the non-zero elements.*/
package main

import (
	"sort"
)

func moveZeroes(arr []int) []int {
	// first value is index, second is value
	sort.Ints(arr)
	zeroCount := 0
	for i, v := range arr {
		if v == 0 {
			zeroCount = i
		} else {
			break
		}
	}
	nonZeros := arr[zeroCount:]
	//zeroCount := arr[:zeroCount]
}
func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}
