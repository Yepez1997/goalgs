package main

import (
	"fmt"
)

// first letter is n // len(alphabet)
// in essence we are figuring how many times to loop to get to the first index
// second letter is n % len(alphabet)
// in other words get the remaining left over words to fill in for the second index
func excelSheet(n int) string {
	var alphabet = "abcdefghijklmnopqrstuvwxyz" // easier if made into a list
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
