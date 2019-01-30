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
		count[sWords[j]] = 1
	}

	for i := 0; i < len(tWords); i++ {
		if _, ok := count[tWords[i]]; ok {
			count[tWords[i]]++
		} else {
			count[tWords[i]] = 1
		}
	}

	results := []string{}
	for key, value := range count {
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
	s := "this apple is sweet"
	t := "this apple is sour"
	result := uncommonWords(s, t)
	fmt.Println(result)
}
