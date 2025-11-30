package main

import "fmt"

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	res := &ListNode{}
	op := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}

		op = op.Next
	}

	if list1 != nil {
		op.Next = list1
	}
	if list2 != nil {
		op.Next = list2
	}

	return res.Next
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	list := mergeTwoLists(list1, list2)

	fmt.Println("list.val: ", list.Val)
	for list != nil {
		list = list.Next
		if list != nil {
			fmt.Println("list.val: ", list.Val)
		}
	}
}
