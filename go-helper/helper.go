package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

// 打印 json
func PrintJSON(t any) {
	jsonstr, _ := json.MarshalIndent(t, "", "  ")
	fmt.Print(string(jsonstr) + "\n")
}

// 生成链表
func GenerateLinkedList(arr []int) ListNode {
	var dummyHead = &ListNode{}
	var curr = dummyHead
	for _, value := range arr {
		var node = &ListNode{value, nil}
		curr.Next = node
		curr = curr.Next
	}
	return *dummyHead.Next
}

// 生成指定长度的随机 int 数组
func GenerateRandomArray(length int, min int, max int) []int {
	arr := make([]int, length)
	if min > max {
		return arr
	}
	idx := 0
	rand.Seed(time.Now().Unix())
	for idx < length {
		arr[idx] = rand.Intn(max - min) + min
		idx++
	}
	return arr
}
