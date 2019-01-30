package main

import (
	"fmt"
)

// work on more this one is fun
// 1 -> A
// 2 -> B
// hash modulo len of the alpabet
// left overs would be the index of the new column
func excelSheet(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	column := ""
	firstChar := alphabet[n%len(alphabet)]

	fmt.Printf("%v", firstChar)
	if n > len(alphabet) {
		// second index is needed

	}
	return column
}

func main() {
	excelSheet(10)
}
