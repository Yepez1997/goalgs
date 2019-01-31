/*Given an array nums, write a function to move all 0's to the end of it
while maintaining the relative order of the non-zero elements.*/
package main

import (
	"sort"
)

// 0 1 0 3 12
// 1 3 12 0 0
// approach is to keep swapping
// if encounter 1 zeros, swap with next 1
// if encounter 2 zeros together, swap zeroes with next 2
// if encounter 3 zeros together, swa[ zeroes with next 3
// the most effective way to implement is:
// to sort and adjust based on slicing and indeing
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
