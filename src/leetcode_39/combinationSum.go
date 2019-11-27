//https://leetcode-cn.com/problems/combination-sum/
/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/
package main

import (
	"fmt"
	"sort"
)

func main(){
	candidates := []int{2,3,6,7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println(result)
}

var res [][]int
var path []int
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	DFS(0, target, candidates)
	re := res
	res = nil
	return re
}

func DFS(start, target int, candidates []int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)  //go语言赋值与拷贝区别  改变源切片的值改变和不改变新切片值
		res = append(res, tmp)
		return
	}

	for i := start; i < len(candidates) && target - candidates[i] >= 0; i++ {
		path = append(path, candidates[i])
		DFS(i, target - candidates[i], candidates)
		path = path[:len(path) - 1]
	}
}
