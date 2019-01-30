package main

import (
	"fmt"
	"strings"
)

func uncommonWords(s string, t string) []string {
	sWords := strings.Split(s, " ")
	tWords := strings.Split(t, " ")

	count := make(map[string]int)
	for j := 0; j < len(sWords); j++ {
		// never makes it in here
		if _, ok := count[sWords[j]]; ok {
			fmt.Println("HERE")
			count[sWords[j]]++

		} else {
			count[sWords[j]] = 1
		}
	}

	for i := 0; i < len(tWords); i++ {
		// never makes it in here
		fmt.Println(tWords[i])
		if _, ok := count[tWords[i]]; ok {
			count[tWords[i]]++

		} else {
			count[tWords[i]] = 1
		}
	}

	results := []string{}
	for key, value := range count {
		fmt.Println(value)
		if value == 1 {
			//fmt.Println(key)
			results = append(results, key)
		}
	}

	return results
}

// create a set ?

func main() {
	// testing words
	//s := "this apple is sweet"
	//t := "this apple is sour"
	a := "apple apple"
	b := "banana"
	//result := uncommonWords(s, t)
	result2 := uncommonWords(a, b)
	//fmt.Println(result)
	fmt.Println(result2)
}
