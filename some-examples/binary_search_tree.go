package main

import (
	"fmt"
	"slices"
)

type (
	Tree struct {
		node *Node
	}
	Node struct {
		Left  *Node
		Right *Node
		Value int
	}
)

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{Value: value}
		return t
	}

	t.node.Insert(value)
	return t
}

func (t *Tree) Print() {
	t.node.Print()
}

func (t *Tree) Find(v int) *Node {
	return t.node.Find(v)
}

func (t *Tree) Remove(value int) {
	if t.node != nil {
		t.node = t.node.Remove(value)
	}
}

func (t *Tree) BreadthFirstSearch() []int {
	list := []int{}
	queue := []*Node{}

	if t.node != nil {
		queue = append(queue, t.node)
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		list = append(list, currentNode.Value)
		queue = slices.Delete(queue, 0, 1)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}

	}

	return list
}

func (t *Tree) BreadthFirstSearchRecursive(queue []*Node, list []int) []int {
	if len(queue) == 0 {
		return list
	}

	currentNode := queue[0]
	list = append(list, currentNode.Value)
	queue = slices.Delete(queue, 0, 1)

	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}

	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}

	return t.BreadthFirstSearchRecursive(queue, list)
}

func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Remove deletes the node with the given value from the subtree rooted at n
// and returns the new root of this subtree. It handles the three standard
// cases: leaf, single child and two children (replace with inorder successor).
func (n *Node) Remove(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Remove(value)
		return n
	}

	if value > n.Value {
		n.Right = n.Right.Remove(value)
		return n
	}

	// found the node to remove
	// case 1: no children
	if n.Left == nil && n.Right == nil {
		return nil
	}

	// case 2: only right child
	if n.Left == nil {
		return n.Right
	}

	// case 3: only left child
	if n.Right == nil {
		return n.Left
	}

	// case 4: two children - find inorder successor (min in right subtree)
	succ := n.Right
	for succ.Left != nil {
		succ = succ.Left
	}

	// replace current node's value with successor's value
	n.Value = succ.Value
	// remove the successor node from right subtree
	n.Right = n.Right.Remove(succ.Value)

	return n
}

func (n *Node) Find(v int) *Node {
	if n == nil {
		return nil
	}

	fmt.Println("\n\n Going through with n.Value: ", n.Value)
	if v == n.Value {
		return n
	}

	if v < n.Value {
		return n.Left.Find(v)
	} else {
		return n.Right.Find(v)
	}
}

func (n *Node) Print() {
	if n == nil {
		return
	}

	fmt.Println(n.Value)
	n.Left.Print()
	n.Right.Print()
}

func main() {
	tree := NewTree()
	//.    10
	//.  7   12
	//. 5   11   25
	//.            30
	tree.
		Insert(10).
		Insert(12).
		Insert(25).
		Insert(11).
		Insert(7).
		Insert(30).
		Insert(5)

	list := tree.BreadthFirstSearchRecursive([]*Node{tree.node}, []int{})
	fmt.Println("\n\n list is: ", list)

	// tree.Print()

	// n := tree.Find(30)
	// fmt.Println("\n\n The n is: ", n.Value)
}
