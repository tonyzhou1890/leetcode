package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{4,2,1,3})
	helper.PrintJSON(insertionSortList(&l1)) // [1,2,3,4]
	l1 = helper.GenerateLinkedList([]int{-1,5,3,4,0})
	helper.PrintJSON(insertionSortList(&l1)) // [-1,0,3,4,5]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Val = -2 << 31
	dummyHead.Next = insertionSortList(head.Next)
	node := dummyHead
	for node.Next != nil {
		if node.Next.Val < head.Val {
			node = node.Next
		} else {
			break
		}
	}
	head.Next = node.Next
	node.Next = head
	return dummyHead.Next
}