// Given an array of integers and an integer k
// you need to find the total number of continuous
// subarrays whose sum equals to k
package main

func subArraysK(arr []int, k int) int {
	// two pointer approach
	left := 0
	subArrayCount := 0
	continousSum := 0
	for right := 0; right < len(arr); right++ {
		// check the first case
		// keep incrementing right
		for continousSum < k && left < right {
			// if we exhausted all possible subarrys
			if left == len(arr) {
				break
			}
			continousSum += arr[left]
		}

		if continousSum == k {
			subArrayCount = subArrayCount + 1
			continousSum = 0
		}
		left++
	}
	return 0
}

// if it is a subarry increase k

// 0 , 1
// 1 , 1
// 1 , 1
// 0 , 1

func main() {
	arr := []int{1, 1, 1}
	target := 2
	subArraysK(arr, k)
	// solution should be two
}

// define the problem
// Lets say there exists some closed points alpha [i .. j]
// such that the sum of i, i + 1, i + 2, i + n, j is k
// Lets define alphaList <- [alpha1, alpha2, ..., alphamm]
// The number of continous subarrays is sum(alphalist)

// define the recursive solution
// sum(i,j - 1) == k
// while i < j - 1
// 		set i++
//

// its a continous approach
// define the the time and space complexity
