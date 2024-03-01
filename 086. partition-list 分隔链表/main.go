package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,4,3,2,5,2})
	helper.PrintJSON(partition(&l1, 3)) // [1,2,2,4,3,5]
	l1 = helper.GenerateLinkedList([]int{2,1})
	helper.PrintJSON(partition(&l1, 2)) // [1,2]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	list1 := &ListNode{}
	list2 := &ListNode{}
	p1 := list1
	p2 := list2
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p1.Next = list2.Next
	p2.Next = nil
	return list1.Next
}