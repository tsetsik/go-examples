package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	seen := map[rune]bool{}
	m := 0
	count := 0
	s = strings.TrimSpace(s)
	for _, c := range s {
		if _, exists := seen[c]; exists {
			m = max(m, count)
			count = 1
			continue
		}
		seen[c] = true
		count++
	}

	return m
}

func main() {
	res := lengthOfLongestSubstring("pwwkew")
	fmt.Println("\n\n res is: ", res)
}
