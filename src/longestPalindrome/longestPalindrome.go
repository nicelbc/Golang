package main

import "fmt"

func main() {
	s := "cbbbbbbd"
	palindrome := longestPalindrome(s)

	fmt.Println(palindrome)

}

func longestPalindrome(str string) string {
	len := len(str)
	if len == 0 || len == 1 {
		return str
	}

	start := 0 //回文串起始位置
	max := 1   //最大长度

	dp := make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, len)
		dp[i][i] = 1
		if i < len-1 && str[i] == str[i+1] {
			dp[i][i+1] = 1
			start = i
			max = 2
		}
	}

	for l := 3; l <= len; l++ { //l表示检索子串的长度，从3开始
		for i := 0; i+l-1 < len; i++ {
			endPose := i + l - 1                                   //终止字符串位置
			if str[i] == str[endPose] && dp[i+1][endPose-1] == 1 { //状态转移
				dp[i][endPose] = 1
				start = i
				max = l
			}
		}
	}

	return str[start : start+max] //返回最长回文子串
}
