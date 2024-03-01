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
// func postorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}
// 	stack := []*TreeNode{}
// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		node := stack[len(stack) - 1]
// 		stack = stack[ : len(stack) - 1]
// 		if node.Left == nil && node.Right == nil {
// 			res = append(res, node.Val)
// 		} else {
// 			stack = append(stack, node)
// 			if node.Right != nil {
// 				stack = append(stack, node.Right)
// 				node.Right = nil
// 			}
// 			if node.Left != nil {
// 				stack = append(stack, node.Left)
// 				node.Left = nil
// 			}
// 		}
// 	}

// 	return res
// }

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]

		if node == nil {
			node := stack[len(stack) - 1]
			stack = stack[ : len(stack) - 1]
			res = append(res, node.Val)
		} else {
			// 用 nil 作为分隔，下一次访问到 nil，说明 nil 之前的节点已经访问过了，可以直接取值
			stack = append(stack, node, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return res
}