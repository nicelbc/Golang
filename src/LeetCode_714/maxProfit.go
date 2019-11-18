//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
	fmt.Println(maxProfitDp(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	vally := prices[0]
	re := 0
	for _, val := range prices {
		if val < vally {
			vally = val
		} else if val-fee > vally {
			re += val - fee - vally
			vally = val - fee
		}
	}

	return re
}

func maxProfitDp(prices []int, fee int) int {
	cash := make([]int, len(prices))
	hold := make([]int, len(prices))
	hold[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		cash[i] = Max(cash[i-1], hold[i-1]+prices[i]-fee)
		hold[i] = Max(hold[i-1], cash[i-1]-prices[i])
	}
	return cash[len(prices)-1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
