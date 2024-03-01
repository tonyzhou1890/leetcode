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
// a 和 b 两个链表长度不一样，但有公共部分 c。所以，pA 走 a 接着走 b，pB 走 b 接着走 a。这样两者长度差就消除了，相遇的地方就是 c 开始的位置。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}