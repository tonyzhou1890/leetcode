package main

import (
	"leetcode/go-helper"
)

func main() {
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 用一个空节点作为标记，每次遍历一个节点的时候，将其 next 指向空节点，再次碰到空节点说明有环
// func hasCycle(head *ListNode) bool {
// 	flag := &ListNode{}
// 	var temp *ListNode
// 	for head != nil {
// 		if head.Next == flag {
// 			return true
// 		}
// 		temp = head
// 		head = head.Next
// 		temp.Next = flag
// 	}
// 	return false
// }

// 快慢指针
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}