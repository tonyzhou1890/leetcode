package main

import (
	// "leetcode/go-helper"
)

func main() {
}

type Node struct {
    Val int
   	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    nodeMap := map[int]*Node{}
	if node == nil {
		return nil
	}

	var gen func(n *Node) *Node
	gen = func(n *Node) *Node {
		if val, ok := nodeMap[n.Val]; ok {
			return val
		}
		nodeMap[n.Val] = &Node{}
		nodeMap[n.Val].Val = n.Val
		if len(n.Neighbors) > 0 {
			for i := 0; i < len(n.Neighbors); i++ {
				nodeMap[n.Val].Neighbors = append(nodeMap[n.Val].Neighbors, gen(n.Neighbors[i]))
			}
		}
		return nodeMap[n.Val]
	}

	gen(node)

	return nodeMap[1]
}