package main

import (
	"leetcode/go-helper"
	"strconv"
	"strings"
	// "strconv"
)

func main() {
	helper.PrintJSON(isAdditiveNumber("112358")) // true
	helper.PrintJSON(isAdditiveNumber("199100199")) // true
	helper.PrintJSON(isAdditiveNumber("123")) // true
	helper.PrintJSON(isAdditiveNumber("101")) // true
	helper.PrintJSON(isAdditiveNumber("347")) // true
	helper.PrintJSON(isAdditiveNumber("000")) // true
	helper.PrintJSON(isAdditiveNumber("8917")) // true
	helper.PrintJSON(isAdditiveNumber("211738")) // true
	helper.PrintJSON(isAdditiveNumber("011235")) // true
	helper.PrintJSON(isAdditiveNumber("12122436")) // true
	helper.PrintJSON(isAdditiveNumber("199111992")) // true
	helper.PrintJSON(isAdditiveNumber("12012122436")) // true
	helper.PrintJSON(isAdditiveNumber("198019823962")) // true
	helper.PrintJSON(isAdditiveNumber("111122335588143")) // true
	helper.PrintJSON(isAdditiveNumber("1980198239625944")) // true
	helper.PrintJSON(isAdditiveNumber("101020305080130210")) // true
	helper.PrintJSON(isAdditiveNumber("199100199299498797")) // true
	helper.PrintJSON(isAdditiveNumber("1123581321345589144")) // true
	helper.PrintJSON(isAdditiveNumber("1002003005008001300")) // true
	helper.PrintJSON(isAdditiveNumber("11111111111011111111111")) // true
}

/**
当第一个和第二个数确定后，其余的数就已经确定了。比如，第二轮相加的数就是第一轮的第二个数和结果。
所以，除了第一轮时第一个数和第二个数需要确定其位数，后面的轮次都只需要检查合法性就行。
因此，枚举第一个数和第二个数，然后检查以这两个数开始的序列是否满足要求。
*/
func isAdditiveNumber(num string) bool {
	length := len(num)
	if length < 3 {
		return false
	}
	/*
	为了简便，这里两轮循环都是与 length 做比较。
	不能与 1/2 length 比较，因为比如 12 位数，分别三位，这就不对了。
	所以，这里可以外层循环与 1/3 length 比较，内层循环与 2/3 length 比较。
	*/
	for addendEnd := 1; addendEnd < length; addendEnd++ {
		for augendEnd := addendEnd + 1; augendEnd < length; augendEnd++ {
			if check(&num, 0, addendEnd, augendEnd, length) {
				return true
			}
		}
	}
	return false
}

func check(num *string, addendStart int, addendEnd int, augendEnd int, length int) bool {
	// 超过一位的数首字母不为 0
	if (addendEnd - addendStart > 1 && (*num)[addendStart] == '0') || (augendEnd - addendEnd > 1 && (*num)[addendEnd] == '0') {
		return false
	}
	sum := stringAdd((*num)[addendStart : addendEnd], (*num)[addendEnd : augendEnd])
	sumLen := len(sum)
	if augendEnd + sumLen <= length && sum == (*num)[augendEnd : augendEnd + sumLen] {
		// 正好结束
		if augendEnd + sumLen == length {
			return true
		}
		// 继续检查
		return check(num, addendEnd, augendEnd, augendEnd + sumLen, length)
	}
	return false
}

/**
字符串相加
*/
func stringAdd(a string, b string) string {
  aLen := len(a)
	bLen := len(b)
	maxLen := aLen
	if bLen > maxLen {
		maxLen = bLen
	}
	maxLen = maxLen + 1
	res := make([]string, maxLen)
	addend := 0
	augend := 0
	sum := 0
	carry := 0
	for i := 0; i < maxLen - 1; i++ {
		addend = 0
		if i < aLen {
			addend, _ = strconv.Atoi(string(a[aLen - i - 1]))
		}
		augend = 0
		if i < bLen {
			augend, _ = strconv.Atoi(string(b[bLen - i - 1]))
		}
		sum = addend + augend + carry
		// helper.PrintJSON(sum)
		if sum > 9 {
			carry = sum / 10
			res[maxLen - i - 1] = strconv.Itoa(sum % 10)
		} else {
			carry= 0
			res[maxLen - i - 1] = strconv.Itoa(sum)
		}
	}
	if carry > 0 {
		res[0] = strconv.Itoa(carry)
	}
	return strings.Join(res, "")
}
