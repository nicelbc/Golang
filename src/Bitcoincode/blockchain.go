package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
)

const dbFile = "blockchain.db" //数据库文件名当前目录下
const blockBucket = "blocks" //名称，



type BlockChain struct {
	tip []byte //二进制数据
	db *bolt.DB //数据库

}

type BlockChainIterator struct {
	currentHash []byte //当前哈希
	db *bolt.DB //数据库
}


//增加一个区块
func (block *BlockChain) AddBlock(data string){

}

//迭代器
func (block *BlockChain) Iterator() *BlockChainIterator{

}

//去的下一个区块
func (it *BlockChainIterator) Next() *Block{

}

//新建一个区块链
func NewBlockChain() *BlockChain{
	var tip []byte //存储二进制数据
	db, err := bolt.Open(dbFile, 0600,nil) //打开数据库
	if err != nil {
		log.Panic(err)//处理数据库打开错误
	}

	//处理数据更新
	err = db.Update(func (tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))//按照名称打开数据表格
		if bucket == nil {
			fmt.Println("当前数据库没有区块链，没有创建一个新的")
			genesis := NewGenesisBlock()//创建创世区块
			bucket, err := tx.CreateBucket([]byte(blockBucket))//创建一个表格
			if err != nil {
				log.Panic(err)//处理创建错误
			}
			err = bucket.Put(genesis.Hash, genesis.Serialize())//存入数据
			if err != nil {
				log.Panic(err)//处理存入错误
			}
			err = bucket.Put([]byte("1"), genesis.Hash)//存入数据
			if err != nil {
				log.Panic(err)//处理存入错误
			}
			tip = genesis.Hash //区块哈希
		} else {
			tip = bucket.Get([]byte("1"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)//处理数据库更新错误
	}
	bc := BlockChain{tip, db}//创建一个区块链
	return  &bc

	}






