//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	l := removeDuplicates(nums)
	fmt.Println(l)
}

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
