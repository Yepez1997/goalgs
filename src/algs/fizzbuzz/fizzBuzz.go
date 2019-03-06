package main

import (
	"fmt"
	"strconv"
)

func reverse(fizzBuzzList []string) []string {
	fizzBuzzListReversed := []string{}
	for j := len(fizzBuzzList) - 1; j > -1; j-- {
		fizzBuzzListReversed = append(fizzBuzzListReversed, fizzBuzzList[j])
	}
	return fizzBuzzListReversed
}

/* LeetCode Problem 412 */
func fizzBuzz(n int) []string {
	fizzBuzzList := []string{}
	fizz := "Fizz"
	buzz := "Buzz"

	for n > 0 {
		if n%3 == 0 && n%5 == 0 {
			fizzBuzzList = append(fizzBuzzList, (fizz + buzz))
			n--
		} else if n%5 == 0 {
			fizzBuzzList = append(fizzBuzzList, (buzz))
			n--
		} else if n%3 == 0 {
			fizzBuzzList = append(fizzBuzzList, (fizz))
			n--
		} else {
			fizzBuzzList = append(fizzBuzzList, strconv.Itoa(n))
			n--
		}
	}
	reversed := reverse(fizzBuzzList)
	return reversed
}

func main() {
	testNum := 15
	result := fizzBuzz(testNum)
	fmt.Println(result)
}
