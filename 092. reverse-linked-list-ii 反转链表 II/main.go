package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{1,2,3,4,5})
	helper.PrintJSON(reverseBetween(&l1, 2, 4)) // [1,4,3,2,5]
	l1 = helper.GenerateLinkedList([]int{5})
	helper.PrintJSON(reverseBetween(&l1, 1, 1)) // [5]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	leftNode := dummyHead

	// 反转部分之前的一个节点
	for i := 0; i < left - 1; i++ {
		leftNode = leftNode.Next
	}

	// 反转部分，从第二个节点开始，每个节点换到区间最前面
	nextPrev := leftNode.Next
	if nextPrev != nil {
		next := nextPrev.Next
		for i := 0; i < right - left; i++ {
			if next != nil {
				nextPrev.Next = next.Next
				next.Next = leftNode.Next
				leftNode.Next = next
				next = nextPrev.Next
			}
		}
	}
	
	return dummyHead.Next
}