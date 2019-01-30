/* author: Aureliano Yepez
* Classic Two Sum Problem
* O(n) time uses a hash/dict to store values
 */

package main

import (
	"sort"
)

//TwoSum -- returns indicies in arr where sum(i,j) == target
func TwoSum(arr []int, target int, res []int) []int {
	// index and val
	hash := make(map[int]int)
	for i, elem := range arr {
		hash[elem] = i
	}

	for j := 0; j < len(arr); j++ {
		complement := target - arr[j]
		// this checks of val := hash[complement]
		if val, ok := hash[complement]; ok {
			res = append(res, j, val)
			sort.Ints(res)
			return res
		}
	}
	return res
}

func main() {

}
