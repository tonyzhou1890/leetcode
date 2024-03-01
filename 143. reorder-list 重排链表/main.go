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
func reorderList(head *ListNode)  {
	// 快慢指针找中点
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 反转后半段
	fast = slow.Next
	slow.Next = nil
	var prev *ListNode
	for fast != nil {
		next := fast.Next
		fast.Next = prev
		prev = fast
		fast = next
	}
	// 拼接链表
	root := head
	for root != nil && prev != nil {
		tempRoot := root.Next
		root.Next = prev
		root = tempRoot
		tempPrev := prev.Next
		prev.Next = root
		prev = tempPrev
	}
}