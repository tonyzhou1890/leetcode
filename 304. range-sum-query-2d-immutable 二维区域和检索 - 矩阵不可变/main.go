package main

import (
	"leetcode/go-helper"
	// "strconv"
)

func main() {
	obj := Constructor([][]int{
		{3,0,1,4,2},
		{5,6,3,2,1},
		{1,2,0,1,5},
		{4,1,0,1,7},
		{1,0,3,0,5},
	})
	helper.PrintJSON(obj)
	helper.PrintJSON(obj.SumRegion(2,1,4,3)) // 8
	helper.PrintJSON(obj.SumRegion(1,1,2,2)) // 11
	helper.PrintJSON(obj.SumRegion(1,2,2,4)) // 12
}


// type NumMatrix struct {
//   matrix [][]int
// 	cache map[string]int
// }


// func Constructor(matrix [][]int) NumMatrix {
// 	var data = new(NumMatrix)
// 	data.matrix = matrix
// 	data.cache = map[string]int{}
// 	return *data
// }


// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
// 	var key = strconv.Itoa(row1) + "-" + strconv.Itoa(col1) + "-" + strconv.Itoa(row2) + "-" + strconv.Itoa(col2)
// 	if val, ok := this.cache[key]; ok {
// 		return val
// 	}
// 	var sum = 0
// 	for y := row1; y <= row2; y++ {
// 		for x := col1; x <= col2; x++ {
// 			sum += this.matrix[y][x]
// 		}
// 	}
// 	this.cache[key] = sum
// 	return sum
// }


/**
前缀和
cache[i][j] 代表 0 行 0 列 到 i 行 j 列的和
*/
type NumMatrix struct {
  matrix [][]int
	cache [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var data = new(NumMatrix)
	data.matrix = matrix
	rows := len(matrix)
	cols := len(matrix[0])
	data.cache = make([][]int, rows + 1)
	for key := range data.cache {
		data.cache[key] = make([]int, cols + 1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			data.cache[i][j] = data.cache[i][j - 1] + data.cache[i - 1][j] - data.cache[i - 1][j - 1] + data.matrix[i - 1][j - 1]
		}
	}
	return *data
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.cache[row2 + 1][col2 + 1] - this.cache[row2 + 1][col1] - this.cache[row1][col2 + 1] + this.cache[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
