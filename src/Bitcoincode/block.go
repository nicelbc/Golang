package main

import (
	"bytes"
	"encoding/gob"
	"github.com/labstack/gommon/log"
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

//对象转化为二进制字节集，可以写入文件
func (block *Block) Serialize() []byte {
	var result bytes.Buffer//开辟内存，存放字节集合
	encoder := gob.NewEncoder(&result)//编码对象创建
	err := encoder.Encode(block)//编码操作
	if err != nil {
		log.Panic(err)//处理错误
	}

	return result.Bytes()//返回字节
}

//读取文件，读到二进制字节集，二进制字节集转化为对象
func DeserializeBlock(data []byte) *Block {
	var block Block//对象存储字节转化的对象
	decoder := gob.NewDecoder(bytes.NewReader(data)) //解码
	err := decoder.Decode(&block)//尝试解码
	if err != nil {
		log.Panic(err)//处理错误
	}

	return &block
}