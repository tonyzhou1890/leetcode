package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
  arr := [6000000]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(255)
	}
	tmp := 0
	l := len(arr) - 5
	startTime := time.Now()
	for i := 0; i < l; i++ {
		tmp = arr[i] + arr[i + 1] + arr[i + 2] + arr[i + 3] + arr[i + 4]
	}

	fmt.Print(tmp, '\n')
	fmt.Print(time.Since(startTime), '\n')
}