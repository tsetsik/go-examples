package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 5, 7, 8, 9}
	s2 := []int{5, 10, 15, 30}
	s := mergeSortedSlices(s1, s2)
	fmt.Println("len: ", len(s), "; s: ", s)
}

func mergeSortedSlices(s1, s2 []int) []int {
	res := []int{}

	i, j := 0, 0

	numLoops := len(s1) + len(s2)
	ij := 0

	for ij < numLoops {
		if i < len(s1) && s1[i] < s2[j] {
			s1Item := s1[i]
			res = append(res, s1Item)
			i++
		} else {
			s2Item := s2[j]
			res = append(res, s2Item)
			j++
		}

		ij++
	}

	return res
}
