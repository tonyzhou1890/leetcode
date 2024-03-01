package main

import (
	"leetcode/go-helper"
)

func main() {
	l1 := helper.GenerateLinkedList([]int{2, 4, 3})
	l2 := helper.GenerateLinkedList([]int{5, 6, 4})
	helper.PrintJSON(addTwoNumbers(&l1, &l2)) // [7, 0, 8]
	l1 = helper.GenerateLinkedList([]int{0})
	l2 = helper.GenerateLinkedList([]int{0})
	helper.PrintJSON(addTwoNumbers(&l1, &l2)) // [0]
	l1 = helper.GenerateLinkedList([]int{9,9,9,9,9,9,9})
	l2 = helper.GenerateLinkedList([]int{9,9,9,9})
	helper.PrintJSON(addTwoNumbers(&l1, &l2)) // [8,9,9,9,0,0,0,1]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) *helper.ListNode {
	dummyHead := &helper.ListNode{}
	prev := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		carry = 0
		node := &helper.ListNode{}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum %= 10

		node.Val = sum
		prev.Next = node
		prev = prev.Next
	}
	return dummyHead.Next
}
