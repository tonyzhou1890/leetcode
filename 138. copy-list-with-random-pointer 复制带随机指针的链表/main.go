package main

import (
	// "leetcode/go-helper"
)

func main() {
}

type Node struct {
    Val int
	Next *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
    nodeMap := map[*Node]*Node{}

	dummyHead := &Node{}
	curr := dummyHead

	for head != nil {
		if v, ok := nodeMap[head]; ok {
			curr.Next = v
		} else {
			curr.Next = &Node{head.Val, nil, nil}
			nodeMap[head] = curr.Next
		}
		curr = curr.Next
		if head.Random != nil {
			if v, ok := nodeMap[head.Random]; ok {
				curr.Random = v
			} else {
				curr.Random = &Node{head.Random.Val, nil, nil}
				nodeMap[head.Random] = curr.Random
			}
		}
		head = head.Next
	}

	return dummyHead.Next
}