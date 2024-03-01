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
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	// 判断子树是否满足数值范围
	var help func(root *TreeNode, min int, max int) bool
	help = func(root *TreeNode, min int, max int) bool {
		if root.Val < min || root.Val > max {
			return false
		}
		// 左子树判断
		if root.Left != nil {
			if help(root.Left, min, root.Val - 1) == false {
				return false
			}
		}
		// 右子树判断
		if root.Right != nil {
			if help(root.Right, root.Val + 1, max) == false {
				return false
			}
		}
		return true
	}

	return help(root, - 2 << 31, 2 << 31 - 1)
}
