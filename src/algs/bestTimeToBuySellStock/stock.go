package main

import "fmt"

// elts of size 2 go program needs fix

// selling price must be larger
func sellStock(prices []int) int {
	fmt.Println("enter")
	runningSum := 0
	maxSum := 0
	for i := 0; i < len(prices); i++ {
		low := prices[i]
		current := i
		fmt.Println(low)
		for low <= prices[current] && current < len(prices)-1 {
			runningSum = (prices[current+1] - low)
			//fmt.Println(prices[current])
			if runningSum > maxSum {
				maxSum = runningSum
			}
			current++
		}
		runningSum = 0
	}
	fmt.Println("exit")
	return maxSum
}

// buy low sell high
func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	profit := sellStock(arr)
	//arr2 := []int{7, 6, 4, 3, 1}
	//noProfit := sellStock(arr2)
	arr3 := []int{1, 2}
	errors := sellStock(arr3)
	 prfmt.Println(profit)
	//fmt.Println(noProfit)
	fmt.Println(errors)
	// [1,2]
}
