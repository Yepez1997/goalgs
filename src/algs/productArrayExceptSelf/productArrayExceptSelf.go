package main

import "fmt"

// works for most cases
// figure cases for zero
// [2,3,5,7]
// 105 35 21 15
func products(nums []int) []int {
	// get first product
	firstProduct := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		} else {
			firstProduct *= nums[i]
		}
	}

	// nums to return
	result := []int{}
	result = append(result, firstProduct)

	// all combinations by the first product
	for j := 1; j < len(nums); j++ {
		if nums[j] == 0 {
			result = append(result, 0)
		} else {
			result = append(result, firstProduct/nums[j])
		}
	}

	return result
}
func main() {
	input1 := []int{2, 3, 5, 7}
	input2 := []int{1, 2, 3, 4}
	output1 := products(input1)
	output2 := products(input2)
	fmt.Println(output1)
	fmt.Println(output2)

}
