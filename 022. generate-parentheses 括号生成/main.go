package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
	helper.PrintJSON(generateParenthesis(1)) // ["()"]
}

// func generateParenthesis(n int) []string {
// 	res := []string{}
// 	if n == 0 {
// 		return res
// 	}
// 	return *join(n, n)
// }

// func join(l int, r int) *[]string {
// 	temp := []string{}
// 	if l == 0 && r == 0 {
// 		return &[]string{""}
// 	}
// 	// 还有左括号
// 	if l > 0 {
// 		for _, s := range *join(l - 1, r) {
// 			temp = append(temp, "(" + s)
// 		}
// 	}
// 	// 剩下的左括号小于右括号，使用右括号
// 	if l < r && r > 0 {
// 		for _, s := range *join(l, r - 1) {
// 			temp = append(temp, ")" + s)
// 		}
// 	}
// 	return &temp
// }

var res = []string{}
func generateParenthesis(n int) []string {
	res = []string{}
	if n == 0 {
		return res
	}
	join("", n, n)
	return res
}

func join(cur string, l int, r int) {
	if l == 0 && r == 1 {
		res = append(res, cur + ")")
		return
	}
	if l > 0 {
		join(cur + "(", l - 1, r)
	}
	if l < r {
		join(cur + ")", l, r - 1)
	}
}