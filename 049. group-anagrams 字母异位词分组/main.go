package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["bat"],["nat","tan"],["ate","eat","tea"]]
	helper.PrintJSON(groupAnagrams([]string{""})) // [[""]]
	helper.PrintJSON(groupAnagrams([]string{"a"})) // [["a"]]
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	cache := map[string][26]int{}
	for _, s := range strs {
		// 当前字符是否有对应分组
		hasGroup := false
		// 建立字符频率表
		freOfS := [26]int{}
		for _, c := range s {
			freOfS[c - 'a'] += 1
		}
		for idx, g := range res {
			// 长度不同，肯定不是同一组
			if len(g[0]) != len(s) {
				continue
			}
			// 比对每一组的第一个字符串
			fre := cache[g[0]]
			// 频率是否相同
			sameFre := true
			for i := 0; i < 26; i++ {
				if freOfS[i] != fre[i] {
					sameFre = false
					break
				}
			}
			if sameFre {
				res[idx] = append(res[idx], s)
				hasGroup = true
				break
			}
		}
		if !hasGroup {
			res = append(res, []string{s})
			cache[s] = freOfS
		}
	}
	return res
}