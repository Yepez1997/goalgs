/* author: Aureliano Yepez
* Classic Two Sum Problem
* O(n) time uses a hash/dict to store values
 */

package main

import (
	"fmt"
	"sort"
)

func twosum(arr []int, target int, res *[]int) {
	// index and val
	hash := make(map[int]int)
	for i, elem := range arr {
		hash[elem] = i
	}

	// check if value is in the hash
	for j := 0; j < len(arr); j++ {
		complement := target - arr[j]
		// this checks of val := hash[complement]
		if val, ok := hash[complement]; ok {
			*res = append(*res, j, val)
			sort.Ints(*res)
			break
		}
	}
}

func main() {
	arr := []int{1, 4, 5, 2, 4}
	target := 7
	res := []int{}
	twosum(arr, target, &res)
	fmt.Printf("%v\n", res)
}
