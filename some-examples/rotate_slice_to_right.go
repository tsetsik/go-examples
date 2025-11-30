package main

import "fmt"

func reverseArr(n []int, start, end int) {
	for start < end {
		n[start], n[end] = n[end], n[start]
		start++
		end--
	}
}

// input: [1,2,3,4,5,6,7]
// output: [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	n := len(nums)
	k = k % n

	reverseArr(nums, 0, n-1)
	reverseArr(nums, 0, k-1)
	reverseArr(nums, k, n-1)
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println("arr: ", arr)
}
