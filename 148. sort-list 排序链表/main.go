package main

import (
	"leetcode/go-helper"
)

func main() {
	arr := []int{1,5,3,4}
	link := helper.GenerateLinkedList(arr)
	sLink := SortList(&link)
	helper.PrintJSON(sLink)
}

func SortList(head *helper.ListNode) *helper.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fast = head
	var lastSlow = head
	var slow = head
	for ;fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		lastSlow = slow
		slow = slow.Next
	}
	var listR = SortList(slow)
	lastSlow.Next = nil
	var listL = SortList(head)
	var newHead = &helper.ListNode{}
	newHead.Val = 0
	var p = newHead

	for ;listR != nil && listL != nil; {
		if listR.Val < listL.Val {
			p.Next = listR
			listR = listR.Next
		} else {
			p.Next = listL
			listL = listL.Next
		}
		p = p.Next
		p.Next = nil
	}
	if listR != nil {
		p.Next = listR
	}
	if listL != nil {
		p.Next = listL
	}
	return newHead.Next
}
