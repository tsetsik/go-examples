package main

import (
	"fmt"
	"slices"
)

// 6, 5, 3
// 5, 6, 3

func insertionSort(in []int) []int {
	slices.Sort(in)
	for i := 1; i < len(in); i++ {
		key := in[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && in[j] > key {
			in[j+1] = in[j]
			j--
		}

		in[j+1] = key
	}

	return in
}

func main() {

	// unsorted := []int{6, 5, 3, 1, 8, 7, 2, 4}
	unsorted := []int{6, 5, 3}
	sorted := insertionSort(unsorted)
	fmt.Println("\n sorted: ", sorted)
}
