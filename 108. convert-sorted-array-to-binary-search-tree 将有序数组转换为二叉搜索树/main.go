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
	helper.PrintJSON(sortedArrayToBST([]int{-10,-3,0,5,9})) // [0,-3,9,-10,null,5]
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var help func(left int, right int) *TreeNode
	help = func(left, right int) *TreeNode {
		idx := (left + right) / 2
		node := &TreeNode{}
		node.Val = nums[idx]
		if idx > left {
			node.Left = help(left, idx - 1)
		}
		if right > idx {
			node.Right = help(idx + 1, right)
		}
		return node
	}

	return help(0, len(nums) - 1)
}