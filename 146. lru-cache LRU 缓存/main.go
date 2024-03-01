package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

type ListNode = helper.ListNode

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1); // 缓存是 {1=1}
	helper.PrintJSON(lRUCache.List.Val)
	lRUCache.Put(2, 2); // 缓存是 {1=1, 2=2}
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	helper.PrintJSON(lRUCache.Get(1));    // 返回 1
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Get(2);    // 返回 -1 (未找到)
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Get(1);    // 返回 -1 (未找到)
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Get(3);    // 返回 3
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
	lRUCache.Get(4);    // 返回 4
	helper.PrintJSON(lRUCache.List.Val)
	helper.PrintJSON(lRUCache.List.Next.Val)
}

type CacheType map[int]*DoubleLinkNode

type DoubleLinkNode struct {
    Key int
	Val int
	Next *DoubleLinkNode
	Prev *DoubleLinkNode
}


type LRUCache struct {
	Cache CacheType
	List *DoubleLinkNode
	Capacity int
	Length int
}


func Constructor(capacity int) LRUCache {
	res := LRUCache{}
    res.Cache = CacheType{}
	res.Capacity = capacity
	return res
}


func (this *LRUCache) Get(key int) int {
	if val, ok := this.Cache[key]; ok {
		if val.Prev != nil {
			head := val
			for head.Prev != nil {
				head = head.Prev
			}
			val.Prev.Next = val.Next
			if val.Next != nil {
				val.Next.Prev = val.Prev
			}
			val.Prev = nil
			val.Next = head
			head.Prev = val
			this.List = val
		}
		return val.Val
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if val, ok := this.Cache[key]; ok {
		val.Val = value
		this.Get(key)
	} else {
		node := &DoubleLinkNode{key, value, nil, nil}
		node.Next = this.List
		if this.List != nil {
			this.List.Prev = node
		}
		this.List = node
		this.Cache[key] = node
		if this.Length == this.Capacity {
			head := this.List
			for head.Next != nil {
				head = head.Next
			}
			head.Prev.Next = nil
			head.Prev = nil
            delete(this.Cache, head.Key)
		} else {
			this.Length++
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */