package main

import (
	"fmt"
)

// P AYP ALIS HIRI NG
// PAHNAPLSIIGYIR
// P   A   H   N
// A P L S I I G
// Y   I   R
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// res := make([][]rune, numRows)

	for _, str := range s {

	}

	return ""
}

func main() {
	str := "PAHNAPLSIIGYIR"
	num := 3
	res := convert(str, num)

	fmt.Println("res: ", res)
}
