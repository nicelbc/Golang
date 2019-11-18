package main

type BlockChain struct {
	blocks [] *Block //一个数组，每个元素都是指针，存储block区块地址
}

//增加一个区块
func (blocks *BlockChain) AddBlock(data string) {
	preBlock := blocks.blocks[len(blocks.blocks) - 1] //去除最后一个区块
	newBlock := NewBlock(data, preBlock.Hash) //创建一个区块
	blocks.blocks = append(blocks.blocks, newBlock)//区块链插入一个新的区块
}

//创建一个区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}