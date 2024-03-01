package main

import (
)

type Node struct {
   Val int
   Left *Node
   Right *Node
   Next *Node
}

func main() {
}


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	stack := []*Node{}
// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		length := len(stack)
// 		for i := 0; i < length; i++ {
// 			node := stack[0]
// 			stack = stack[1 : length]
// 			node.Next = nil
// 			if i < length - 1 {
// 				node.Next = stack[0]
// 			}
// 			if node.Left != nil {
// 				stack = append(stack, node.Left, node.Right)
// 			}
// 		}
// 	}

// 	return root
// }

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	start := root
	curr := root
	prev := root

	for curr != nil {
		// start 设为下一行的起始
		if start == curr {
			start = curr.Left
			prev = start
		}
		if curr.Left != nil {
			prev.Next = curr.Left
			curr.Left.Next = curr.Right
			prev = curr.Right
		}
		if curr.Next != nil {
			curr = curr.Next
		} else {
			curr = start
		}
	}

	return root
}