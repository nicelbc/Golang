
//https://leetcode-cn.com/problems/minimum-falling-path-sum/
package main

import "fmt"

func main() {
	A := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	re := minFallingPathSum(A)
	fmt.Println(re)
}

func minFallingPathSum(A [][]int) int {
	rows := len(A)
	cols := rows

	for row := rows-2; row >=0; row-- {
		for col := 0; col < cols; col++ {
			min := A[row+1][col]
			if col > 0 {
				min = Min(min, A[row+1][col-1])
			}
			if col+1 < cols {
				min = Min(min, A[row+1][col+1])
			}
			A[row][col] += min
		}
	}

	re := A[0][0]
	for _, val := range A[0] {
		re = Min(val, re)
	}
	return re
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}