package main

import (
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

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
func recoverTree(root *TreeNode)  {
	var nodeX, nodeY, prep *TreeNode

	stack := []*TreeNode{}

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		if prep != nil && root.Val < prep.Val {
			nodeY = root
			if nodeX == nil {
				nodeX = prep
			} else {
				break
			}
		}
		prep = root
		root = root.Right
	}

	nodeX.Val, nodeY.Val = nodeY.Val, nodeX.Val
}
