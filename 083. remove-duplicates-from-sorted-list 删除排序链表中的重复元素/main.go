package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,1,2})
	helper.PrintJSON(deleteDuplicates(&l1)) // [1,2]
	l1 = helper.GenerateLinkedList([]int{1,1,2,3,3})
	helper.PrintJSON(deleteDuplicates(&l1)) // [1,2,3]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}