package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2})) // 3
	helper.PrintJSON(canCompleteCircuit([]int{2,3,4}, []int{3,4,3})) // -1
}

// 从 i 开始，检查之后每一步的结余是否大于等于 0
// func canCompleteCircuit(gas []int, cost []int) int {
// 	length := len(gas)

// 	for i := 0; i < length; i++ {
// 		balance := 0
// 		for steps := 0; steps < length; steps++ {
// 			balance += gas[(i + steps) % length] - cost[(i + steps) % length]
// 			// 无法到达下一站，则从下一站开始。假设从 x 经过 y 到不了 z，则从 x 到 z 之间的任一站都无法到达 z，因为经过之前任一站的时候汽油肯定是大于等于 0 的。
// 			if balance < 0 {
// 				i += steps
// 				break
// 			}
// 		}
// 		if balance >= 0 {
// 			return i
// 		}
// 	}

// 	return -1
// }

// 假设，从 y 可以到达终点，结余油量 m。y 点之前存在数个亏空，总亏空为 n。那么，只要 m 大于等于 n，即可以弥补第一站到 y 的所有亏空，从 y 出发就可以走完全程。
// 亏空 n 为负数，结余 m 为非负数，n + m >= 0 即可以走完全程，即从第一站到最后一站，油总量大于消耗即可。
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)

	balance := 0
	debt := 0
	start := 0

	for i := 0; i < length; i++ {
		balance += gas[i] - cost[i]
		debt += gas[i] - cost[i]
		if balance < 0 {
			balance = 0
			start = i + 1
		}
	}

	if debt >= 0 {
		return start
	}
	return -1
}