package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(evalRPN([]string{"2","1","+","3","*"})) // 9
	helper.PrintJSON(evalRPN([]string{"4","13","5","/","+"})) // 6
	helper.PrintJSON(evalRPN([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"})) // 22
}

// 栈操作
func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		l := len(stack)
		var val int
		if tokens[i] == "+" {
			val = stack[l - 2] + stack[l - 1]
		} else if tokens[i] == "-" {
			val = stack[l - 2] - stack[l - 1]
		} else if tokens[i] == "*" {
			val = stack[l - 2] * stack[l - 1]
		} else if tokens[i] == "/" {
			val = stack[l - 2] / stack[l - 1]
		} else {
			val, _ := strconv.Atoi(string(tokens[i]))
			stack = append(stack, val)
			continue
		}
		stack = stack[ : l - 2]
		stack = append(stack, val)
	}
	return stack[0]
}