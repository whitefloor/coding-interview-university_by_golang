package main

import (
	"errors"
	"fmt"
	"log"
)

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

// ArrayBoundaryCheck 傳回該所引的元素值並附帶邊界檢查
func ArrayBoundaryCheck(arr []int, index int) (int, error) {
	if index < 0 || index > len(arr)-1 {
		return 0, errors.New("index out of array range")
	}

	return (arr)[index], nil
}

// ArrayInsertElement 在任意位置插入元素
func ArrayInsertElement(arr []int, element, index int) {
	if len(arr) == index { // nil or empty slice or after last element
		arr = append(arr, element)
	}
	arr = append(arr[:index+1], arr[index:]...) // index < len(a)
	arr[index] = element

	fmt.Println(arr)
}

// ArrayPopLastElement 從陣列取出並刪除最後一個元素,回傳刪除過後的陣列
func ArrayPopLastElement(arr []int) (popedArray []int, popValue int) {
	popValue = arr[len(arr)-1]
	popedArray = arr[:len(arr)-1]

	return
}

// ArrayDeleteIndexElement 從陣列中刪除該所引的值,並回傳陣列
func ArrayDeleteIndexElement(arr []int, index int) []int {
	switch {
	case index == 0:
		return arr[index+1:]
	case index == len(arr)-1:
		return arr[:index]
	case index > 0 && index < len(arr)-1:
		arr = append(arr[:index], arr[index+1:]...)
	default:
		log.Println("index out of range can not remove element")
		return nil
	}

	return nil
}

// ArrayRemoveTheSameValue 從陣列中找到相同的值並刪除
func ArrayRemoveTheSameValue(arr []int, value int) []int {
	newArr := make([]int, 0)
	for _, v := range arr {
		if v != value {
			newArr = append(newArr, v)
		}
	}

	return newArr
}

// ArrayReturnTheValueIndex 找到第一個符合value的並回傳索引,找不到就回傳-1
func ArrayReturnTheValueIndex(arr []int, value int) (index int) {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}

// 沒有實作的,覺得這不實際 ======

// resize(nex_capacity) // private function
// 當陣列已經用盡了所有容量後，把陣列的容量*2。
// 如果移除掉一個元素後，陣列實際大小是最大容量的1/4，則把陣列容量減半。
