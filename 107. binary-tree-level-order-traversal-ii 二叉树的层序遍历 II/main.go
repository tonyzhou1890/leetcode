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
func levelOrderBottom(root *TreeNode) [][]int {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	res := [][]int{}
	for len(stack) != 0 {
		tempStack := []*TreeNode{}
		row := []int{}
		for i := 0; i < len(stack); i++ {
			row = append(row, stack[i].Val)
			if stack[i].Left != nil {
				tempStack = append(tempStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				tempStack = append(tempStack, stack[i].Right)
			}
		}
		res = append(res, row)
		stack = tempStack
	}
	for i, length := 0, len(res); i < length / 2; i++ {
		res[i], res[length - i - 1] = res[length - i - 1], res[i]
	}
	return res
}