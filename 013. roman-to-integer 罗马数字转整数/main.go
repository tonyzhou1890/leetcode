package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(romanToInt("III")) // 3
	helper.PrintJSON(romanToInt("IV")) // 4
	helper.PrintJSON(romanToInt("IX")) // 9
	helper.PrintJSON(romanToInt("LVIII")) // 58
	helper.PrintJSON(romanToInt("MCMXCIV")) // 1994
}

type RomanMap struct {
	roman string
	val int
}

// func romanToInt(s string) int {
// 	table := []RomanMap{
// 		{"I",1},
// 		{"IV",4},
// 		{"V",5},
// 		{"IX",9},
// 		{"X",10},
// 		{"XL",40},
// 		{"L",50},
// 		{"XC",90},
// 		{"C",100},
// 		{"CD",400},
// 		{"D",500},
// 		{"CM",900},
// 		{"M",1000},
// 	}
	
// 	num := 0
// 	for i := len(table) - 1; s != ""; i-- {
// 		romanLen := len(table[i].roman)
// 		for len(s) >= romanLen && s[0:romanLen] == table[i].roman {
// 			s = s[romanLen:]
// 			num += table[i].val
// 		}
// 	}
// 	return num
// }

func romanToInt(s string) int {
	table := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	
	num := 0
	for i := 0; i < len(s); i++ {
		if i < len(s) - 1 && table[s[i]] < table[s[i + 1]] {
			num += table[s[i + 1]] - table[s[i]]
			i++
		} else {
			num += table[s[i]]
		}
	}
	return num
}