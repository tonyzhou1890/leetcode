package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,3,3,4,4,5})
	helper.PrintJSON(deleteDuplicates(&l1)) // [1,2,5]
	l1 = helper.GenerateLinkedList([]int{1,1,1,2,3})
	helper.PrintJSON(deleteDuplicates(&l1)) // [2,3]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	dummyHead := &ListNode{}
// 	dummyHead.Next = head
// 	prev := dummyHead
// 	next := dummyHead.Next
	
// 	for next != nil {
// 		if next.Next != nil && next.Val == next.Next.Val {
// 			// 直到想通的值只存在一个
// 			for next.Next != nil && next.Val == next.Next.Val {
// 				next.Next = next.Next.Next
// 			}
// 			// 去掉相同的值
// 			prev.Next = next.Next
// 			next = prev.Next
// 		} else {
// 			prev = next
// 			next = next.Next
// 		}
// 	}
// 	return dummyHead.Next
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	
	// 头重复--不要头
	if head.Next != nil && head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		return deleteDuplicates(head.Next)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}