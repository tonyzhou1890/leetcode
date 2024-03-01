package main

import (
	"leetcode/go-helper"
)

func main() {
	// arr := []int{1, 2, 3, 5};
	arr := []int{0,1,3,5,6};
	// arr := []int{100};
	// arr := []int{100,100,100};
	// arr := []int{0,0,0};
	helper.PrintJSON(hIndex(arr))
}

// 套百度的公式 要确定一个人的h指数非常容易，到SCI网站，查出某个人发表的所有SCI论文，让其按被引次数从高到低排列，往下核对，直到某篇论文的序号大于该论文被引次数
// 二分查找
func hIndex(citations []int) int {
	left := 0
	length := len(citations)
	right := length - 1
	/**
	left 就是不满足条件的部分的第一个元素。
	要想 left 变成大于 right，只能是 left 等于 right 这一轮的结果。
	而 left 等于 right 则是 right 比 left 大 1 的那一轮的结果。在 right 比 left 大 1 的轮次，mid 等于 left。left 处于满足条件的一侧往不满足条件的一侧走。所以只能是 left + 1 得到 left 等于 right 的轮次。在 left 等于 right 的轮次。left 是不满足条件的，因为 left 步入了 right 的范围。这时 right - 1，跳出循环。

	那为什么第一个不满足的 left 就是结果，而不是 left - 1 呢。毕竟这是个突变点。
	因为，left - 1 的文章数大于引用数，h 是 citations[left - 1]，left 的文章数则等于或者小于引用数。left 的文章数比 left - 1 的文章数小 1，而 citations[left - 1] 比如 left - 1 的文章数最起码也小 1。所以 left 的文章数是大于等于 citations[left - 1] 的。
	*/
	for ;left <= right; {
		mid := (right - left) / 2 + left
		if (length - mid) > citations[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return length - left
}