// Given an array of integers and an integer k
// you need to find the total number of continuous
// subarrays whose sum equals to k
package main

import "fmt"

// another implementation is to use a hashmap
func subArraysK(arr []int, k int) int {
	cache := make(map[int]int)
	numSubArrs := 0
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == k {
			numSubArrs++
		}
		if v, ok := cache[sum-k]; ok {
			numSubArrs += v
		}
		if v, ok := cache[sum]; !ok {
			cache[sum] = 1
		} else {
			cache[sum] = v + 1
		}
	}
	return numSubArrs
}

// if it is a subarry increase k
// [1,1,1], 2
// 0 , 1
// 1 , 1
// 1 , 1
// 0 , 1

func main() {
	arr := []int{1, 1, 1}
	target := 2
	res := subArraysK(arr, target)
	fmt.Println(res)
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
