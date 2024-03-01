package main

import (
	"leetcode/go-helper"
	"math"
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
func isBalanced(root *TreeNode) bool {
	var help func(root *TreeNode) (depth int, valid bool)
	help = func(root *TreeNode) (depth int, valid bool) {
		if root == nil {
			return 0, true
		}
		leftDpeth, leftValid := help(root.Left)
		rightDepth, rightValid := help(root.Right)

		depth = leftDpeth
		if rightDepth > leftDpeth {
			depth = rightDepth
		}
		depth += 1

		return depth, leftValid && rightValid && math.Abs(float64(leftDpeth - rightDepth)) <= 1
	}

	_, valid := help(root)
	return valid
}