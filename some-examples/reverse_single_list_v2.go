package main

import "fmt"

type (
	node struct {
		value int
		next  *node
	}

	linkedList struct {
		head   *node
		tail   *node
		length int
	}
)

// 1 -> 2 -> 3 -> 4 -> 5
// 2 -> 3
// 2 -> 1; first = second; second = third
// 5 -> 4 -> 3 -> 2 -> 1
func (ll *linkedList) reverse() {
	first := ll.head
	firstUnmodified := first

	second := first.next

	for second != nil {
		third := second.next

		second.next = first
		first = second

		second = third
	}

	ll.head = first

	firstUnmodified.next = nil
	ll.tail = firstUnmodified
}

func (ll *linkedList) prepend(value int) {
	ll.length++

	n := &node{value: value}
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}

	n.next = ll.head
	ll.head = n
}

func (ll *linkedList) append(value int) {
	ll.length++
	n := &node{value: value}
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}

	ll.tail.next = n
	ll.tail = n
}

func (ll *linkedList) print() {
	if ll.head == nil {
		return
	}

	fmt.Println("\nhead.value: ", ll.head.value)

	current := ll.head
	nodeIndex := 1
	for current != nil && nodeIndex < (ll.length-1) {
		current = current.next
		if current == nil {
			continue
		}
		fmt.Println("\nnode.value: ", current.value, " at index: ", nodeIndex)
		nodeIndex++
	}

	fmt.Println("\ntail.value: ", ll.tail.value)
}

func main() {
	ll := linkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)

	ll.reverse()

	ll.print()
}
