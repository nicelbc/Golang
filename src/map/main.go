package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}

	//统计每个数字出现的次数
	countMap := make(map[int]int)
	for _, val := range arr1 {
		if _, ok := countMap[val]; ok {
			countMap[val]++
		} else {
			countMap[val] = 1
		}
	}
	fmt.Println("map中的存储顺序 key : value\n", countMap)

	fmt.Println("遍历输出map中key value")
	for key, val := range countMap {
		fmt.Printf("key: %d value: %d\n", key, val)
	}
	fmt.Println("再次遍历输出map中key value")
	for key, val := range countMap {
		fmt.Printf("key: %d value: %d\n", key, val)
	}

	//排序输出
	//定义切片数组存放key
	keys := []int{}
	for key, _ := range countMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	fmt.Println("排序输出:")
	for _, key := range keys {
		fmt.Printf("key: %d value: %d\n", key, countMap[key])
	}

}

/*
	总结： map遍历取值的时候 是无序的
		   如果需要按顺序取map的key, 需要另外定义数组排序sort()
*/
