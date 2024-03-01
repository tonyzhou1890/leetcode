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

	var start, curr, prev *Node
	curr = root
	prev = &Node{}

	for curr != nil {
		// start 设为下一行的起始
		if start == nil {
			if curr.Left != nil {
				start = curr.Left
			} else if curr.Right != nil {
				start = curr.Right
			}
		}
		
		if curr.Left != nil {
			prev.Next = curr.Left
			prev = prev.Next
		}
		if curr.Right != nil {
			prev.Next = curr.Right
			prev = prev.Next
		}

		if curr.Next != nil {
			curr = curr.Next
		} else {
			curr = start
			start = nil
			prev = &Node{}
		}
	}

	return root
}