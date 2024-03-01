package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	// arr := []int{1, 2, 3, 5};
	// arr := []int{3,0,6,1,5};
	arr := []int{100};
	helper.PrintJSON(hIndex(arr))
}

// func hIndex(citations []int) int {
// 	h := 0

// 	for _, value := range citations {
// 		greaterThanValueCount := 0
// 		for _, value2 := range citations {
// 			if value2 >= value {
// 				greaterThanValueCount++
// 			}
// 		}
// 		tempH := value
// 		if greaterThanValueCount < tempH {
// 			tempH = greaterThanValueCount
// 		}
// 		if tempH > h {
// 			h = tempH
// 		}
// 	}

// 	return h
// }
// 套百度的公式 要确定一个人的h指数非常容易，到SCI网站，查出某个人发表的所有SCI论文，让其按被引次数从高到低排列，往下核对，直到某篇论文的序号大于该论文被引次数
func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)));
	h := 0;
	for key, value := range citations {
		// 文章数超过引用次数，结束
		if ((key + 1) > value) {
			break;
		}
		h = key + 1;
	}
	return h
}