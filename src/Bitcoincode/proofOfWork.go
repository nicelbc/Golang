package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64 //最大的64位整数
)
const targetBits = 24 //对比的位数

type ProofOfWork struct {
	block *Block //区块
	target *big.Int //存储特定哈希值的整数
}

//创建一个工作量证明的挖矿对象
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1) //初始化目标整数
	target.Lsh(target,uint(256-targetBits))//数据转换
	pow := &ProofOfWork{block,target}
	return pow
}

//准备数据挖矿
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,//上一块哈希
			pow.block.Data,//当前数据
			IntToHex(pow.block.Timestamp),//时间转为十六进制
			IntToHex(int64(targetBits)),//位数
			IntToHex(int64(nonce)),//保存工作量的nonce
		}, []byte{},
		)

	return data
}

//挖矿
func (pow *ProofOfWork) Run() (int, []byte ){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("当前挖矿计算的数据%s",pow.block.Data)
	for nonce < maxNonce {
		data := pow.PrepareData(nonce)//准备好的数据
		hash = sha256.Sum256(data)//计算哈希
		fmt.Printf("\r%x",hash)//打印显示哈希
		hashInt.SetBytes(hash[:])//获取要对比的数据
		if hashInt.Cmp(pow.target) == -1 {//挖矿校验
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n\n")
	return nonce, hash[:]// nonce解题的答案， 当前哈希
}

//校验挖矿成功
func (pow *ProofOfWork) Validate()bool {
	var hashInt big.Int
	data := pow.PrepareData(pow.block.Nonce)//准备好的数据
	hash := sha256.Sum256(data)//计算哈希
	hashInt.SetBytes(hash[:])//获取要对比的数据
	isValid := (hashInt.Cmp(pow.target) == -1)//校验数据
	return isValid
}