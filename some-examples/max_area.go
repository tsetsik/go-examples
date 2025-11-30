package main

import "fmt"

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0

	for i < j {
		minH := min(height[i], height[j])
		area := (j - i) * minH

		res = max(res, area)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(height)
	fmt.Println("res is: ", res)
}
