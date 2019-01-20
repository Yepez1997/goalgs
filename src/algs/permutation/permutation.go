package main

import (
	"fmt"
	"sort"
	"strings"
)

/* Given two strings find if there exists a permutation */
func charList(s string) string {
	sSplit := strings.Split(s, "")
	sort.Strings(sSplit)
	return strings.Join(sSplit, "")

}

func permuation(s string, t string) bool {
	if len(s) != len(t) || (charList(s) != charList(t)) {
		return false
	}
	return true
}
func main() {
	oneStr := "ahjklsw"
	twoStr := "ahjlews"
	res := permuation(oneStr, twoStr)
	s := fmt.Sprintf("a %v", res)
	fmt.Println(s)
}
