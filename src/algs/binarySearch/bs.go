package main

import "fmt"

// BinarySearch -- returns value in O(log(n))
func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2 // finds the middle
		if arr[mid] == target {
			return arr[mid]
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	target := 4
	sortedarr := []int{1, 2, 4, 5, 8, 9, 12}
	result := BinarySearch(sortedarr, target)
	fmt.Printf("%d\n", result)
}
