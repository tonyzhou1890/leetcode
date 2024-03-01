package main

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	res := [][]int{}
	level := 0
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
		level++
		if level % 2 == 0 {
			for i, length := 0, len(row); i < length / 2; i++ {
				row[i], row[length - i - 1] = row[length - i - 1], row[i]
			}
		}
		res = append(res, row)
		stack = tempStack
	}
	return res
}