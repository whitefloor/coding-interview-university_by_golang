package main

import (
	"fmt"
	"log"
)

func main() {
	times := BinarySearch(75)
	log.Println(times)
}

// 補充：線性搜尋法(Linear search)
// 就是從頭開始找，時間複雜度O(n)

// BinarySearch 二分搜尋法-搜尋已排序過的數列
func BinarySearch(target int) (times int) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}

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
