package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(intToRoman(3)) // "III"
	helper.PrintJSON(intToRoman(4)) // "IV"
	helper.PrintJSON(intToRoman(9)) // "IX"
	helper.PrintJSON(intToRoman(58)) // "LVIII"
	helper.PrintJSON(intToRoman(1994)) // "MCMXCIV"
}

type RomanMap struct {
	roman string
	val int
}

func intToRoman(num int) string {
	table := []RomanMap{
		{"I",1},
		{"IV",4},
		{"V",5},
		{"IX",9},
		{"X",10},
		{"XL",40},
		{"L",50},
		{"XC",90},
		{"C",100},
		{"CD",400},
		{"D",500},
		{"CM",900},
		{"M",1000},
	}
	
	str := ""
	for i := len(table) - 1; num > 0; i-- {
		if num >= table[i].val {
			times := num / table[i].val
			num -= times * table[i].val
			for times > 0 {
				str += table[i].roman
				times--
			}
		}
	}
	return str
}
