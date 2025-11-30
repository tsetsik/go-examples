package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var res strings.Builder
	valueSlice := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	keySlice := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// check till number becomes 0
	for num > 0 {
		// check valueSlice in given order with number
		for idx, val := range valueSlice {
			// if number is greater than current element in valueSlice, reduce that value from number and append that character in final-string. and break from inner loop to check number with valueSlice again from 1000
			if num >= val {
				res.WriteString(keySlice[idx])
				num = num - val
				break
			}
		}
	}

	return res.String()
}

func main() {
	romanStr := intToRoman(64)
	fmt.Println("romanStr: ", romanStr)
}
