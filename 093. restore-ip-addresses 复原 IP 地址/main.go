package main

import (
	"leetcode/go-helper"
	"strings"
)

func main() {
	helper.PrintJSON(restoreIpAddresses("25525511135")) // ["255.255.11.135","255.255.111.35"]
	helper.PrintJSON(restoreIpAddresses("0000")) // ["0.0.0.0"]
	helper.PrintJSON(restoreIpAddresses("101023")) // ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
}

func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}
	var help func(idx int)
	help = func(idx int) {
		// 正好分为四块
		if len(path) == 4 && idx >= len(s) {
			res = append(res, strings.Join(path, "."))
			return
		}
		if idx >= len(s) {
			return
		}
		// 1 位
		path = append(path, s[idx : idx + 1])
		help(idx + 1)
		path = path[ : len(path) - 1]
		// 2 位
		if idx < len(s) - 1 && s[idx] != '0' {
			path = append(path, s[idx : idx + 2])
			help(idx + 2)
			path = path[ : len(path) - 1]
		}
		// 3 位
		if idx < len(s) - 2 && ((s[idx] == '1') || (s[idx] == '2' && ((s[idx + 1] < '5') || (s[idx + 1] == '5' && s[idx + 2] <= '5')))) {
			path = append(path, s[idx : idx + 3])
			help(idx + 3)
			path = path[ : len(path) - 1]
		}
	}

	help(0)

	return res
}