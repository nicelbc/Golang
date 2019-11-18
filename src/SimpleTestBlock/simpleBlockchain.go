package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo"
	"time"
)

//区块模型
type BlockModel struct {
	Id        int64  //索引
	Timestamp string //时间戳
	BPM       int    //心跳频率每分钟
	Hash      string //哈希值 sha256
	PreHash   string //上一块哈希值
}

//区块链创建，数组
var BlockChain = make([]BlockModel, 0)

//创世模块，第一个区块
func init() {
	//创建一个区块
	block := BlockModel{0, time.Now().String(), 0, "", ""}
	BlockChain = append(BlockChain, block) // 加入数组
}

//哈希处理
func calcHash(block BlockModel) string {
	record := string(block.Id) + block.Timestamp + string(block.BPM) + block.PreHash
	myhash := sha256.New()            //创建sha256对象
	myhash.Write([]byte(record))      //加入数据
	hashed := myhash.Sum(nil)         //计算哈希
	return hex.EncodeToString(hashed) //编码为字符串
}

//区块有效性检验
func Is_BlockValid(newBlock, lastBlock BlockModel) bool {
	//id不能相等，必须是顺序的
	if lastBlock.Id+1 != newBlock.Id {
		return false
	}
	//前个区块哈希，不等于新区块上一个哈希
	if lastBlock.Hash != newBlock.PreHash {
		return false
	}
	if calcHash(newBlock) != newBlock.Hash { //数据被修改
		return false
	}

	return true
}

//处理区块的创建
func createBlock(ctx echo.Context) error {
	//处理心跳信息
	type message struct {
		BPM int
	}
	var mymessage = message{}
	if err := ctx.Bind(&mymessage); err != nil { //绑定消息
		panic(err) //处理错误
	}

	lastblock := BlockChain[len(BlockChain)-1] //前一个区块
	//使用前一个区块创建新区块
	newblock := BlockModel{}
	newblock.Id = lastblock.Id + 1           //序列号+1
	newblock.Timestamp = time.Now().String() //当前时间
	newblock.BPM = mymessage.BPM
	newblock.PreHash = lastblock.Hash
	newblock.Hash = calcHash(newblock)
	if Is_BlockValid(newblock, lastblock) {
		BlockChain = append(BlockChain, newblock)
		fmt.Println("创建区块成功", BlockChain[len(BlockChain)-1].Id)
	} else {
		fmt.Println("创建区块失败")
	}

	return ctx.JSON(200, newblock)
}

func main() {
	echoserver := echo.New() //创建服务器
	echoserver.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, BlockChain)
	})
	echoserver.POST("/", createBlock)
	echoserver.Logger.Fatal(echoserver.Start(":8848"))
}
