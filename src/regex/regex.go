package main

import (
	"fmt"
	"regexp"
)

const text = "My email is mail@gmail.com@abc.com" +
	"  email1 is abc@de.org  email2 is   kk@qq.com  " +
	" email3 id ddd@abc.com.cn   email4 is ddd@abc.com.cn"

func main() {
	re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\\.[a-zA-Z0-9\\.]+)")

	//匹配一个
	match := re.FindString(text)
	fmt.Println(match)

	//匹配所有符合条件的
	matchall := re.FindAllString(text, -1)
	fmt.Println(matchall)

	//子匹配
	matchsub := re.FindAllStringSubmatch(text, -1)
	for _, m := range matchsub {
		fmt.Println(m)
	}
}

/*
. 匹配任意字符
+ 匹配一次或者多次
* 匹配零次或多次
\\. 转义字符匹配符号点（.）
() 加括号 用于匹配后的子匹配...submatch()
*/
