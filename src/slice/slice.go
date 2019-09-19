package main

import "fmt"

func main() {
	slice := make([]int, 10, 20) //定义slice的 len 10 cap 20
	fmt.Printf("%p \n", slice) // slice 指针地址
	fmt.Printf("%d %d\n", len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("%p \n", slice) // 没超过cap 地址不变
	fmt.Printf("%d %d \n", len(slice), cap(slice))

	slice = append(slice, 6, 7, 8, 9, 10, 11)
	fmt.Printf("%p \n", slice) // 超过cap 地址改变 并且cap翻倍
	fmt.Printf("%d %d \n", len(slice), cap(slice))
}


/*
切片slice的len和cap
容量表示底层数组的大小，长度是你可以使用的大小。
当你用 append扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉。也就是说，容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。
而长度，则是为了帮助你限制切片可用成员的数量，提供边界查询的。所以用 make 申请好空间后，需要注意不要越界【越 len 】
扩容,容量成倍数进行扩容
*/
