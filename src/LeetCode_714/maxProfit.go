//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit(prices, fee))
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
