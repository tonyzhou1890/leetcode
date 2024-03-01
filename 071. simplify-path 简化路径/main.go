package main

import (
	"leetcode/go-helper"
	"strings"
)

func main() {
	helper.PrintJSON(simplifyPath("/home/")) // "/home"
	helper.PrintJSON(simplifyPath("/../")) // "/"
	helper.PrintJSON(simplifyPath("/home//foo/")) // "/home/foo"
	helper.PrintJSON(simplifyPath("/a/./b/../../c/")) // "/c"
}

func simplifyPath(path string) string {
	res := []string{}
	pathArr := strings.Split(path, "/")
	for i := 0; i < len(pathArr); i++ {
		// . 忽略
		if pathArr[i] == "." || pathArr[i] == "" {
			continue
		}
		// 返回上一级
		if pathArr[i] == ".." {
			if len(res) > 0 {
				res = res[ : len(res) - 1]
			}
			continue
		}
		// 有效目录
		res = append(res, pathArr[i])
	}
	return "/" + strings.Join(res, "/")
}