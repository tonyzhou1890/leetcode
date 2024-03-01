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
func buildTree(inorder []int, postorder []int) *TreeNode {
	hashMap := map[int]int{}
	for idx, val := range inorder {
		hashMap[val] = idx
	}

	var help func(inorderLeft int, inorderRight int, postorderLeft int, postorderRight int) *TreeNode
	help = func(inorderLeft int, inorderRight int, postorderLeft int, postorderRight int) *TreeNode {
		root := &TreeNode{}
		root.Val = postorder[postorderRight]
		inorderRootIdx := hashMap[root.Val]
		if inorderLeft < inorderRootIdx {
			root.Left = help(inorderLeft, inorderRootIdx - 1, postorderLeft, postorderLeft + (inorderRootIdx - inorderLeft) - 1)
		}
		if inorderRootIdx < inorderRight {
			root.Right = help(inorderRootIdx + 1, inorderRight, postorderLeft + (inorderRootIdx - inorderLeft), postorderRight - 1)
		}
		return root
	}

	if len(inorder) == 0 {
		return nil
	}
	return help(0, len(inorder) - 1, 0, len(postorder) - 1)
}