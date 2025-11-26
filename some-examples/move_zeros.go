package main

import "fmt"

// [0,1,0,3,12]
// [1,3,12,0,0]
func moveZeroes(nums []int) {
	store := 0
	for i, n := range nums {
		if n != 0 {
			nums[store], nums[i] = nums[i], nums[store]
			store++
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println("\n\n arr is: ", arr)
}
