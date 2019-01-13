package main

import "fmt"

// would be awesome if go supported templates
/* basic binary search on sorted array */
func bs(arr []int, target int, res *int) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + (high - low)) / 2 // finds the middle
		if arr[mid] == target {
			fmt.Println("here")
			*res = arr[mid]
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

func main() {
	res := -1
	target := 4
	sortedarr := []int{1, 2, 4, 5, 8, 9, 12}
	bs(sortedarr, target, &res)
	// print result
	fmt.Printf("%v\n", res) // shouldn return 4

}
