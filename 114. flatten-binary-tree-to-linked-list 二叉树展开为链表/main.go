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
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	stack := []*TreeNode{}
	if root == nil {
		return
	}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		node.Left = nil
		node.Right = nil
		if len(stack) > 0 {
			node.Right = stack[len(stack) - 1]
		}
	}
}