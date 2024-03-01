package test

import (
	"leetcode/go-helper"
)

func test() {
	arr := []any{1.2,2,3,4}
	link := helper.GenerateLinkedList(arr)
	helper.PrintJSON(link)
}

