package main

import (
	"leetcode/go-helper"
)

func main() {
	// arr := []int{2,2,2,2,2,2};
	arr := []int{1,3,4,2,2};
	helper.PrintJSON(findDuplicate(arr))
}

/**
快慢指针
将元素值作为下标，可以形成一个链表的访问路径。因为存在相同的元素，则形成了环，因此可以通过快慢指针找到重复数。
1. 快慢指针必然在环中相遇，并且这时慢指针并没有走完环的一圈。设快指针已经走完了 n 圈，这时距离满指针一步，则下一步相遇。如果快指针距离满指针两步，因为每次接近慢指针一步，则两步后相遇。即快指针每次缩短与慢指针一步的距离，快指针不可能跳过慢指针，其必然在慢指针一圈未走完时相遇，并且这时快指针的步数是慢指针的两倍。
2. 设起点距离环入口 l 步，环长 m 步。相遇时，慢指针距离环入口 a 步。快指针已经走了 n（n >= 1）圈。则有：
  l + n * m + a = 2 * (l + a)
=> l = n * m - a
这时，让快指针回到起点，与慢指针一同一步一步走，其必然在环入口相遇。
*/
// func findDuplicate(nums []int) int {
// 	fast := 0;
// 	slow := 0;
// 	num := 0
// 	for true {
// 		fast = nums[nums[fast]];
// 		slow = nums[slow]
// 		if slow == fast {
// 			fast = 0;
// 			for nums[fast] != nums[slow] {
// 				fast = nums[fast];
// 				slow = nums[slow];
// 			}
// 			num = nums[fast];
// 			break
// 		}
// 	}
// 	return num
// }

/**
位图法
因为数值小于 n，所以可以开辟一个长度为 n + 1 的数组，然后标记这个数是否出现过，如果碰到出现过的数，说明重复了。
这种方式需要 O(n) 空间。
*/
func findDuplicate(nums []int) int {
	length := len(nums)
  tags := make([]bool, length);
	num := 0;
	for _, value := range nums {
		if tags[value] == true {
			num = value;
			break;
		} else {
			tags[value] = true
		}
	}
	return num
}