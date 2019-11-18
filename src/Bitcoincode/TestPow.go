package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func mainx() {
	start := time.Now()
	for i := 0; i < 10000000000; i++ { //循环挖矿
		data := sha256.Sum256([]byte(strconv.Itoa(i))) //计算哈希
		fmt.Printf("%10d,%x\n",i,data)
		fmt.Printf("%s\n",string(data[len(data)-2:]))
		if string(data[len(data)-2:]) == "00" { //哈希匹配
			usedtime := time.Since(start) //消耗时间
			fmt.Println("Mining success",usedtime)
			break
		}
	}
}