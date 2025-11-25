package main

import (
	"fmt"
)

func selectionSort(in []int) []int {
	for j := 0; j < len(in); j++ {
		lowIndex := j
		lowestNumber := in[j]
		for i := j; i < len(in); i++ {
			current := in[i]

			if current < lowestNumber {
				lowestNumber = current
				lowIndex = i
			}
		}

		tmpLow := in[lowIndex]
		tmpTo := in[j]

		in[j] = tmpLow
		in[lowIndex] = tmpTo
	}

	return in
}

func main() {
	foo := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	// foo := []int{3, 2, 1}
	res := selectionSort(foo)
	fmt.Println("\n\n res is: ", res)
}
