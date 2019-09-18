//https://leetcode-cn.com/problems/relative-sort-array/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	relativeSortArray(arr1, arr2)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m_map := make(map[int]int)
	for _, val := range arr1 {
		if _, ok := m_map[val]; ok {
			m_map[val]++
		} else {
			m_map[val] = 1
		}
	}

	re := []int{}
	for _, val := range arr2 {
		re = Append(re, val, m_map[val])
		delete(m_map, val)
	}

	notin := []int{}
	for key, _ := range m_map {
		notin = append(notin, key)
	}
	sort.Ints(notin)
	for _, val := range notin {
		re = Append(re, val, m_map[val])
	}
	fmt.Println(re)
	return re
}

func Append(sli []int, val int, count int) []int {
	for i := 0; i < count; i++ {
		sli = append(sli, val)
	}
	return sli
}
