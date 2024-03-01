package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(getHint("1807", "7810")) // 1A3B
	helper.PrintJSON(getHint("011", "110")) // 1A2B
	helper.PrintJSON(getHint("1122", "1222")) // "3A0B"
}

// func getHint(secret string, guess string) string {
// 	length := len(secret)
// 	flags := make([]bool, length)
// 	bulls := 0
// 	cows := 0
// 	for i := 0; i < length; i++ {
// 		// 判断是否相同--bull
// 		if guess[i] == secret[i] {
// 			bulls++
// 			flags[i] = true
// 			continue
// 		}
// 		// 不相同，则寻找是否存在需要调整位置的对应字符--cow
// 		for j := 0; j < length; j++ {
// 			// 已经匹配过的跳过
// 			if flags[j] == true {
// 				continue
// 			}
// 			if secret[j] == guess[i] && secret[j] != guess[j] {
// 				flags[j] = true
// 				cows++
// 				break
// 			}
// 		}
// 	}
// 	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
// }

// func getHint(secret string, guess string) string {
// 	length := len(secret)
// 	// 标记数字的抵消情况
// 	flags := make([]int, 10)
// 	bulls := 0
// 	cows := 0
// 	for i := 0; i < length; i++ {
// 		// 判断是否相同--bull
// 		if guess[i] == secret[i] {
// 			bulls++
// 			continue
// 		}
// 		// 不相同--cow
// 		gidx, _ := strconv.Atoi(string(guess[i]))
// 		// 如果有可以用来抵消的数字，计数
// 		if flags[gidx] > 0 {
// 			cows++
// 		}
// 		// 库存减一
// 		flags[gidx]--
// 		// 存入当前可抵消数字
// 		sidx, _ := strconv.Atoi(string(secret[i]))
// 		// 如果库存亏空，计数
// 		if flags[sidx] < 0 {
// 			cows++
// 		}
// 		flags[sidx]++
// 	}
// 	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
// }

func getHint(secret string, guess string) string {
	length := len(secret)
	// 标记数字的抵消情况
	flags := make([]int, 10)
	bulls := 0
	cows := 0
	for i := 0; i < length; i++ {
		// 判断是否相同--bull
		if guess[i] == secret[i] {
			bulls++
			continue
		}
		// 不相同--cow
		// 如果有可以用来抵消的数字，计数
		if flags[guess[i] - '0'] > 0 {
			cows++
		}
		// 库存减一
		flags[guess[i] - '0']--
		// 存入当前可抵消数字
		// 如果库存亏空，计数
		if flags[secret[i] - '0'] < 0 {
			cows++
		}
		flags[secret[i] - '0']++
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
