package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,3,4,5})
	helper.PrintJSON(reverseList(&l1)) // [5,4,3,2,1]
	l1 = helper.GenerateLinkedList([]int{1,2})
	helper.PrintJSON(reverseList(&l1)) // [2,1]
	helper.PrintJSON(reverseList(nil)) // []
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	prev := &ListNode{}
// 	prev = nil
// 	curr := head
// 	if head == nil {
// 		return head
// 	}
// 	next := head.Next
// 	for curr != nil {
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 		if next != nil {
// 			next = next.Next
// 		}
// 	}
// 	return prev
// }

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}