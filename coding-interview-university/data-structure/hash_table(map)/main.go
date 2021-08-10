package main

import "fmt"

func main() {
	mp := make(map[int]int)

	// add key and value
	for i := 0; i < 10; i++ {
		mp[i] = i
	}

	// 遍歷(Traversal)map的方法,取得key及value
	for key, value := range mp {
		fmt.Println(key, value)
	}

	// update key-value and check the key is exists
	if _, ok := mp[1]; ok {
		mp[1] = 100
		fmt.Println(mp[1])
	}

	// delete
	delete(mp, 1)
	if _, ok := mp[1]; ok {
		mp[1] = 100
		fmt.Println(mp[1])
	} else {
		fmt.Println("key doesn't not exists")
	}
}
