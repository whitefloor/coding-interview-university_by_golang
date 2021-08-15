package main

import (
	"fmt"
	"log"
)

func main() {
	v := BinarySearch(75)
	log.Println(v)
}

// BinarySearch 搜尋已排序過的數列
func BinarySearch(target int) int {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}

	times := 0
	LeftSide := 0
	RightSide := len(arr) - 1

	for LeftSide < RightSide {
		times++
		now := (RightSide-LeftSide)/2 + LeftSide // 等於(L+R)/2，但用非註解的部分可以避免overflow
		fmt.Println(arr[now], now)
		if arr[now] == target {
			return times
		} else if arr[now] > target {
			RightSide = now
		} else if arr[now] < target {
			LeftSide = now
		}
	}

	return 0
}
