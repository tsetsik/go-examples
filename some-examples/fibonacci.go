package main

import "fmt"

func fibonacciIterrative(n int) int {
	if n < 2 {
		return n
	}

	currentNum := 0
	nextNum := 1
	for i := 1; i < n; i++ {
		tmp := nextNum + currentNum
		currentNum = nextNum
		nextNum = tmp
	}

	return nextNum
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func main() {
	num := 8

	fibIterative := fibonacciIterrative(num)
	fmt.Println("\n fibonacciIterrative: ", fibIterative)

	fibRecursive := fibonacciRecursive(num)
	fmt.Println("\n fibRecursive: ", fibRecursive)
}
