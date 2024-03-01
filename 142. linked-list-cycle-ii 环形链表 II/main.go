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
func detectCycle(head *ListNode) *ListNode {
    fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			anotherSlow := head
			for slow != anotherSlow {
				slow = slow.Next
				anotherSlow = anotherSlow.Next
			}
			return slow
		}
	}

	return nil
}