package main

import (
	"leetcode/go-helper"
)

func main() {
	l := helper.GenerateLinkedList([]int{1,2,3,4,5})
	helper.PrintJSON(removeNthFromEnd(&l, 2)) // [1,2,3,5]
	l = helper.GenerateLinkedList([]int{1})
	helper.PrintJSON(removeNthFromEnd(&l, 1)) // []
	l = helper.GenerateLinkedList([]int{1,2})
	helper.PrintJSON(removeNthFromEnd(&l, 1)) // [1]
}

type ListNode = helper.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, second, dummyHead := &ListNode{}, &ListNode{}, &ListNode{}
	dummyHead.Next = head
	second = dummyHead
	first = head
	// 快指针先走 n 步，然后同步走，当快指针到达末尾（走完），慢指针的下一个就是要删除的节点
	for n > 0 {
		if first != nil {
			first = first.Next
		}
		n--
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	if second != nil && second.Next != nil {
		second.Next = second.Next.Next
	}
	return dummyHead.Next
}