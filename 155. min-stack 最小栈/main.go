package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

type ListNode = helper.ListNode

func main() {
}

type StackElement struct {
	Val int
	Min int // 截止到当前元素的最小值
}

type MinStack struct {
	Stack []*StackElement
}


func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	min := val
	if len(this.Stack) > 0 {
		if this.Stack[len(this.Stack) - 1].Min < min {
			min = this.Stack[len(this.Stack) - 1].Min
		}
	}
	this.Stack = append(this.Stack, &StackElement{val, min})
}


func (this *MinStack) Pop()  {
	this.Stack = this.Stack[ : len(this.Stack) - 1]
}


func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack) - 1].Val
}


func (this *MinStack) GetMin() int {
	return this.Stack[len(this.Stack) - 1].Min
}

