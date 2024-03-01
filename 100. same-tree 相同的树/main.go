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
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	stackP := []*TreeNode{}
// 	stackQ := []*TreeNode{}

// 	res := true

// 	for len(stackP) != 0 || len(stackQ) != 0 || p != nil || q != nil {
// 		for p != nil {
// 			stackP = append(stackP, p)
// 			p = p.Left
// 		}
// 		for q != nil {
// 			stackQ = append(stackQ, q)
// 			q = q.Left
// 		}
// 		if len(stackP) != len(stackQ) {
// 			res = false
// 			break
// 		}
// 		p = stackP[len(stackP) - 1]
// 		stackP = stackP[ : len(stackP) - 1]
// 		q = stackQ[len(stackQ) - 1]
// 		stackQ = stackQ[ : len(stackQ) - 1]
// 		if p.Val != q.Val {
// 			res = false
// 			break
// 		}
// 		p = p.Right
// 		q = q.Right
// 	}
// 	return res
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && p == q {
		return true
	}
	if (p != q && (p == nil || q == nil)) || p.Val != q.Val {
		return false
	}
	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}