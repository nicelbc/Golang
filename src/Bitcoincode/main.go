package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Game start!")
	bc := NewBlockChain() //创建区块链
	bc.AddBlock("A pay B 10")
	bc.AddBlock("B pay C 13")
	bc.AddBlock("C pay D 31")

	for _, block := range bc.blocks {
		fmt.Printf("上一块哈希%x\n",block.PreBlockHash)
		fmt.Printf("数据：%s\n",block.Data)
		fmt.Printf("当前哈希%x\n",block.Hash)
		pow := NewProofOfWork(block)//检验工作量
		fmt.Printf("pow %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
