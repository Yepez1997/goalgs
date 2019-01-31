package main

import (
	"fmt"
)

// work on more this one is fun
// 1 -> A
// 2 -> B
// hash modulo len of the alpabet
// left overs would be the index of the new column
// divide is the number of times in the list
// that we had to go through for the first
// the modulo is the remainder for the second letter
func excelSheet(n int) string {
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	firstChar := alphabet[int(n/len(alphabet))]

	if firstChar > 0 {
		secondChar := alphabet[int(n%len(alphabet))]
		column := firstChar + secondChar
		return string(column)
	}

	return string(firstChar)
}

func main() {
	strOne := excelSheet(10)
	//strTwo := excelSheet(28)
	//strThree := excelSheet(701)
	fmt.Print(strOne)
	//fmt.Print(strTwo)
	//fmt.Print(strThree)
	// str three is buggy
}
