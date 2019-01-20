package main

import (
	"fmt"
	"sort"
	"strings"
)

/* Given two strings find if there exists a permutation */

/* sortS sorts the string s*/
func sortS(s string) string {
	sSplit := strings.Split(s, "")
	sort.Strings(sSplit)
	return strings.Join(sSplit, "")

}

/* permutation returns turns if there exist a permutation */
func permuation(s string, t string) bool {
	if len(s) != len(t) || (sortS(s) != sortS(t)) {
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
