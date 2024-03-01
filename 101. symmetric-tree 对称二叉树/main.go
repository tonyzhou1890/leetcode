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
func isSymmetric(root *TreeNode) bool {
	var help func(l *TreeNode, r *TreeNode) bool
	help = func(l *TreeNode, r *TreeNode) bool {
		if l == nil && l == r {
			return true
		}
		if (l == nil || r == nil) && l != r {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		return help(l.Left, r.Right) && help(l.Right, r.Left)
	}
	return help(root, root)
}