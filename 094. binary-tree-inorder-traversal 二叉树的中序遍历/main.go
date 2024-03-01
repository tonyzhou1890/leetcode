package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

func main() {
	tree := &TreeNode{}
	tree.Val = 1
	tree.Left = nil
	tree.Right = &TreeNode{}
	tree.Right.Val = 2
	tree.Right.Left = &TreeNode{}
	tree.Right.Left.Val = 3
	helper.PrintJSON(inorderTraversal(tree))
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	res := []int{}
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]

		if node.Right == nil && node.Left == nil {
			res = append(res, node.Val)
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
		}

		stack = append(stack, node)

		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return res
}