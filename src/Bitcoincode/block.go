package main

import (
	"time"
)

//定义区块
type Block struct {
	Timestamp int64 //时间戳，1970年01月01日00.00.00
	Data []byte //交易数据
	PreBlockHash []byte //上一块数据哈希
	Hash []byte //当前哈希
	Nonce int //工作量证明
}

/*//设定结构体对象的哈希
func (block *Block) SetHash() {
	//处理当前时间，转换为十进制字符串，在转换为字节集
	timestamp := []byte(strconv.FormatInt(block.Timestamp,10))
	//叠加要哈希数据
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp},[]byte{})
	//计算哈希值
	hash := sha256.Sum256(headers)
	block.Hash = hash[:] //设置哈希
}*/
//创建一个区块
func NewBlock(data string, preBlockHash []byte) *Block {
	//block是一个指针，取得对象初始化之后的地址
	block := &Block{
		time.Now().Unix(),[]byte(data),preBlockHash,[]byte{},0}
	pow := NewProofOfWork(block)//挖矿附加这个区块
	nonce, hash := pow.Run()//开始挖矿
	block.Hash = hash[:]
	block.Nonce = nonce


	//block.SetHash()//设置哈希
	return block
}
//创建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("创世区块",[]byte{})
}