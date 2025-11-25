package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int64{2, 2, 1, 4, 5, 6}

	res := binarySearchExists(arr, 2)
	fmt.Println(res)

}

func binarySearchExists(array []int64, to_search int64) bool {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	found := false
	low := 0
	high := len(array) - 1
	for low <= high {
		fmt.Println("(low + high) / 2: (", low, "+", high, ") / 2 = ", (low+high)/2)
		mid := (low + high) / 2
		if array[mid] == to_search {
			found = true
			break
		}

		if array[mid] > to_search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}
