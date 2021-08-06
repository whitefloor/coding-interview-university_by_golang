package main

import "fmt"

func main() {
	// 動態生成陣列,使用代數進行初始化
	lens := 10
	arr := make([]int, lens)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println("arr:", arr)

	// 如果想要擷取slice並宣告成新的slice,這是錯誤方法,原因是slice的資料結構會參照到同一個pointer
	errorSlice := arr[2:8]
	fmt.Println("New errorSlice:", errorSlice)

	arr[2] = 100
	fmt.Println("errorSlice value:", errorSlice)

	// 正確做法
	arr[2] = 2
	rightSlice := make([]int, len(arr[2:8]))
	copy(rightSlice, arr[2:8])
	fmt.Println("rightSlice:", rightSlice)

	// 陣列長度,只包含可見的元素
	slice := arr[2:8]
	sliceLen := len(slice)
	fmt.Println("arrLen:", sliceLen)

	// 陣列容量,包含所有元素,平常還沒實際運用過
	sliceCap := cap(slice)
	fmt.Println("arrCap:", sliceCap)

	// 看到所有元素的方法,平常還沒實際運用過
	fmt.Println(slice[:cap(slice)])

	// 清空陣列的方法
	arr = make([]int, lens)
	fmt.Println("arr:", arr)
}
