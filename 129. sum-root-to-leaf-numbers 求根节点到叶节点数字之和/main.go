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
// 层序遍历，更新每个节点的值，叶子节点值相加
func sumNumbers(root *TreeNode) int {
	total := 0
	stack := []*TreeNode{}
	if root == nil {
		return total
	}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1 : ]
		// 叶子节点
		if node.Left == nil && node.Right == nil {
			total += node.Val
		} else {
			// 更新子节点
			if node.Left != nil {
				node.Left.Val += node.Val * 10
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				node.Right.Val += node.Val * 10
				stack = append(stack, node.Right)
			}
		}
	}

	return total
}