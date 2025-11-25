package main

import (
	"fmt"
)

type (
	pairs [][2]int
)

func main() {
	ints := []int{1, 2, 3, 4, 5}

	b := hasPairWithSum(ints, 8)

	fmt.Println("\n\n hasPairWithSum: ", b)
}

func hasPairWithSum(in []int, sum int) bool {
	if len(in) == 0 {
		return false
	}

	seen := map[int]bool{}
	for _, s := range in {
		seen[s] = true
	}

	for _, n := range in {
		numCheck := sum - n
		if seen[numCheck] {
			return true
		}
	}

	return false
}
