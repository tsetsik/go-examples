package main

import "fmt"

type (
	Node struct {
		Value int
		Next  *Node
	}

	LinkedList struct {
		Head   *Node
		Tail   *Node
		Length int
	}
)

func main() {
	ll := LinkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)

	ll.reverse()

	ll.print()
}

func (ll *LinkedList) print() {
	fmt.Println("\n head.value: ", ll.Head.Value, " at index: 0")

	current := ll.Head
	for i := 1; i < ll.Length-1; i++ {
		current = current.Next
		if current == nil {
			continue
		}

		fmt.Println("\n node.value: ", current.Value, " at index: ", i)
	}

	fmt.Println("\n tail.value: ", ll.Tail.Value, " at index: ", ll.Length-1)
}

func (ll *LinkedList) append(value int) {
	n := &Node{Value: value}

	if ll.Tail == nil {
		ll.Tail = n
		ll.Head = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}

	ll.Length++
}

func (ll *LinkedList) prepend(value int) {
	n := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		n.Next = ll.Head
		ll.Head = n
	}

	ll.Length++
}

func (ll *LinkedList) reverse() {
	prev := ll.Head
	ll.Tail = prev

	next := prev

	//
	for next != nil {
		tmp := next.Next

		next.Next = prev
		prev = next

		next = tmp
	}

	ll.Head = prev
	ll.Tail.Next = nil
}

func (ll *LinkedList) insert(index, value int) {
	n := &Node{Value: value}

	if index >= ll.Length {
		ll.append(value)
		return
	}

	if index == 0 {
		ll.prepend(value)
		return
	}

	var prev *Node
	current := ll.Head
	for i := 1; i < index; i++ {
		current = current.Next
		if i == index-1 {
			prev = current
		}
	}

	n.Next = prev.Next
	prev.Next = n

	ll.Length++
}
