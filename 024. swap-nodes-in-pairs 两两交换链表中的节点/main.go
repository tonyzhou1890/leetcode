package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,3,4})
	helper.PrintJSON(swapPairs(&l1)) // [2,1,4,3]
	helper.PrintJSON(swapPairs(nil)) // []
	l1 = helper.GenerateLinkedList([]int{1})
	helper.PrintJSON(swapPairs(&l1)) // [1]
	l1 = helper.GenerateLinkedList([]int{1,2,3,4,5})
	helper.PrintJSON(swapPairs(&l1)) // [2,1,4,3,5]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	first := &ListNode{}
	second := &ListNode{}
	p := dummyHead
	for p != nil {
		if p.Next != nil && p.Next.Next != nil {
			first = p.Next
			second = first.Next
			temp := second.Next
			p.Next = second
			second.Next = first
			first.Next = temp
			p = first
		} else {
			p = nil
		}
	}
	return dummyHead.Next
}