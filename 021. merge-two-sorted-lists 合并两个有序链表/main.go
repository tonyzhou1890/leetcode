package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,4})
	l2 := helper.GenerateLinkedList([]int{1,3,4})
	helper.PrintJSON(mergeTwoLists(&l1, &l2)) // [1,1,2,3,4,4]
	helper.PrintJSON(mergeTwoLists(nil, nil)) // []
	l2 = helper.GenerateLinkedList([]int{0})
	helper.PrintJSON(mergeTwoLists(nil, &l2)) // [0]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	res.Val = 0
	res.Next = nil
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
		p.Next = nil
	}
	if list1 != nil {
		p.Next = list1
	} else if list2 != nil {
		p.Next = list2
	}
	return res.Next
}