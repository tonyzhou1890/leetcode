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
func buildTree(preorder []int, inorder []int) *TreeNode {
	hashMap := map[int]int{}
	for idx, val := range inorder {
		hashMap[val] = idx
	}

	var help func(preorderLeft int, preorderRight int, inorderLeft int, inorderRight int) *TreeNode
	help = func(preorderLeft int, preorderRight int, inorderLeft int, inorderRight int) *TreeNode {
		root := &TreeNode{}
		root.Val = preorder[preorderLeft]
		inorderRootIdx := hashMap[root.Val]
		if inorderLeft < inorderRootIdx {
			root.Left = help(preorderLeft + 1, preorderLeft + (inorderRootIdx - inorderLeft), inorderLeft, inorderRootIdx - 1)
		}
		if inorderRootIdx < inorderRight {
			root.Right = help(preorderLeft + (inorderRootIdx - inorderLeft) + 1, preorderRight, inorderRootIdx + 1, inorderRight)
		}
		return root
	}

	if len(preorder) == 0 {
		return nil
	}
	return help(0, len(preorder) - 1, 0, len(inorder) - 1)
}