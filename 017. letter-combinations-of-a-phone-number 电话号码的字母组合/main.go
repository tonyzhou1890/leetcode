package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	helper.PrintJSON(letterCombinations("")) // []
	helper.PrintJSON(letterCombinations("2")) // ["a","b","c"]
}

func letterCombinations(digits string) []string {
	hash := map[rune][]string{
		'2': {"a","b","c"},
		'3': {"d","e","f"},
		'4': {"g","h","i"},
		'5': {"j","k","l"},
		'6': {"m","n","o"},
		'7': {"p","q","r", "s"},
		'8': {"t","u","v"},
		'9': {"w","x","y","z"},
	}
	res := []string{}
	temp := []string{""}
	for _, num := range digits {
		strs, ok := hash[num]
		if ok {
			res = temp
			temp = []string{}
			for _, exist := range res {
				for _, newChar := range strs {
					temp = append(temp, exist + newChar)
				}
			}
			res = temp
		}
	}
	return res
}