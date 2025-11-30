package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	maxSum, sum := nums[0], 0

	for _, num := range nums {
		sum += num
		maxSum = max(maxSum, sum)
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := maxSubArray(nums)
	fmt.Println("output is: ", output)
}
