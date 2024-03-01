package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

type ListNode = helper.ListNode

func main() {
	list := helper.GenerateLinkedList([]int{-10,-3,0,5,9})
	helper.PrintJSON(sortedListToBST(&list)) // [0,-3,9,-10,null,5]
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sortedListToBST(head *ListNode) *TreeNode {
// 	nums := []int{}
// 	for head != nil {
// 		nums = append(nums, head.Val)
// 		head = head.Next
// 	}

// 	var help func(left int, right int) *TreeNode
// 	help = func(left, right int) *TreeNode {
// 		idx := (left + right) / 2
// 		node := &TreeNode{}
// 		node.Val = nums[idx]
// 		if idx > left {
// 			node.Left = help(left, idx - 1)
// 		}
// 		if right > idx {
// 			node.Right = help(idx + 1, right)
// 		}
// 		return node
// 	}

// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	return help(0, len(nums) - 1)
// }

func sortedListToBST(head *ListNode) *TreeNode {
	p := head
	length := 0
	for p != nil {
		length++
		p = p.Next
	}

	var help func(left int, right int) *TreeNode
	help = func(left, right int) *TreeNode {
		idx := (left + right) / 2
		node := &TreeNode{}
		
		if idx > left {
			node.Left = help(left, idx - 1)
		}

		node.Val = head.Val
		head = head.Next

		if right > idx {
			node.Right = help(idx + 1, right)
		}
		return node
	}

	if length == 0 {
		return nil
	}

	return help(0, length - 1)
}