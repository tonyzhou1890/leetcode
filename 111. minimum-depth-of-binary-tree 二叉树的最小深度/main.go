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
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	leftDepth := minDepth(root.Left)
// 	rightDepth := minDepth(root.Right)

// 	// 叶子节点
// 	if leftDepth == 0 && rightDepth == 0 {
// 		return 1
// 	}

// 	// 没有左节点
// 	if leftDepth == 0 {
// 		return rightDepth + 1
// 	}

// 	// 没有右节点
// 	if rightDepth == 0 {
// 		return leftDepth + 1
// 	}

// 	if leftDepth < rightDepth {
// 		return leftDepth + 1
// 	}

// 	return rightDepth + 1
// }

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	min := 0
	level := 1
	for min == 0 {
		newStack := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			if stack[i].Left == nil && stack[i].Right == nil {
				min = level
				break
			}
			if stack[i].Left != nil {
				newStack = append(newStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				newStack = append(newStack, stack[i].Right)
			}
		}
		stack = newStack
		level++
	}
	return min
}