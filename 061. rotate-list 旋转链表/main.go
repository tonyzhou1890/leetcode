package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,3,4,5})
	helper.PrintJSON(rotateRight(&l1, 2)) // [4,5,1,2,3]
	l1 = helper.GenerateLinkedList([]int{0,1,2})
	helper.PrintJSON(rotateRight(&l1, 4)) // [2,0,1]
	l1 = helper.GenerateLinkedList([]int{1,2})
	helper.PrintJSON(rotateRight(&l1, 0)) // [1,2]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	end := head
	count := 0
	for p != nil {
		count++
		end = p
		p = p.Next
	}
	p = head
	k %= count
	if k == 0 {
		return head
	}
	k = count - k
	pre := head
	for k > 0 {
		pre = p
		p = p.Next
		k--
	}
	end.Next = head
	pre.Next = nil
	return p
}