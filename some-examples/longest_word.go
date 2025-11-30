package main

import (
	"fmt"
	"regexp"
)

func LongestWord(sen string) string {
	r := regexp.MustCompile("[A-Za-z0-9]+")
	strs := r.FindAllString(sen, -1)
	l := map[int]string{}
	m := 0

	for _, s := range strs {
		n := len(s)
		if _, exists := l[n]; exists {
			continue
		}

		l[n] = s
		m = max(m, n)
	}

	if str, exists := l[m]; exists {
		return str
	}

	// code goes here
	return sen
}

func main() {
	fmt.Println(LongestWord("fun&!! time"))
}
