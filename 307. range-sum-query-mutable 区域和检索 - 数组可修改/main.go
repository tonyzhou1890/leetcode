package main

import (
	"leetcode/go-helper"
	"math"
	// "strconv"
)

func main() {
	numArray := constructor([]int{1, 3, 5})
	helper.PrintJSON(numArray.SumRange(0, 2)) // 返回 1 + 3 + 5 = 9
	numArray.Update(1, 2);   // nums = [1,2,5]
	helper.PrintJSON(numArray.SumRange(0, 2)) // 返回 1 + 2 + 5 = 8
}

/**
分块处理。
*/
type NumArray struct {
  size int
	nums []int
	sums []int
}


func constructor(nums []int) NumArray {
	res := NumArray{nums: nums}
	length := len(nums)
	res.size = int(math.Sqrt(float64(length)))
	res.sums = make([]int, (length + res.size - 1) / res.size)
	for key, val := range nums {
		res.sums[key / res.size] += val
	}
	return res
}


func (this *NumArray) Update(index int, val int)  {
	diff := val - this.nums[index]
	this.nums[index] = val
	this.sums[index / this.size] += diff
}


func (this *NumArray) SumRange(left int, right int) int {
	leftBlock := left / this.size
	leftBlockEnd := leftBlock * this.size + this.size
	rightBlock := right / this.size
  rightBlockStart := rightBlock * this.size
	sum := 0
	// 在同一个区块，直接 left 到 right 相加
	if leftBlock == rightBlock {
		for i := left; i <= right; i++ {
			sum += this.nums[i]
		}
	} else {
		// 左区块
		for i := left; i < leftBlockEnd; i++ {
			sum += this.nums[i]
		}
		// 中间区块
		for i := leftBlock + 1; i < rightBlock; i++ {
			sum += this.sums[i]
		}
		// 右区块
		for i := rightBlockStart; i <= right; i++ {
			sum += this.nums[i]
		}
	}
	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */