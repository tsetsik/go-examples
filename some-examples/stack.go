package main

import "fmt"

type (
	stackNode struct {
		value int
		next  *stackNode
	}
	stack struct {
		top    *stackNode
		bottom *stackNode
		length int
	}
)

func (s *stack) peek() *stackNode {
	return s.top
}

func (s *stack) push(value int) *stack {
	sNode := &stackNode{value: value}
	if s.top == nil {
		s.top = sNode
		s.bottom = sNode
	} else {
		sNode.next = s.top
		s.top = sNode
	}

	s.length++
	return s
}

func (s *stack) print() {
	fmt.Println("\n top: ", s.top.value, " at index: 0")

	current := s.top.next
	for i := 1; i < s.length-1; i++ {
		current = current.next
		fmt.Println("\n stackNode: ", current.value, " at index: ", i)
	}

	fmt.Println("\n bottom: ", s.bottom.value, " at index: ", s.length-1)
}

func (s *stack) pop() *stackNode {
	return nil
}

func main() {
	s := &stack{}
	s.push(1).
		push(2).
		push(3)

	s.print()
}
