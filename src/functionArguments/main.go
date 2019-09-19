package main

import "fmt"

func main() {
	//array := make([]int, 0)
	array := []int{}
	array = append(array, 23)
	fmt.Printf("before update:%v\n", array)

	update(&array)
	fmt.Printf("after update:%v\n", array)

}
//函数参数是slice 并且改变slice的值
//定义形参 指向slice的指针
func update(array *[]int) {
	//*array = append(*array, 1)

}

/*
总结：
	golang中的切片slice底层通过数组实现，slice类似一个结构体，其中一个字段保存的是底层数组的地址，
还有长度(len) 和 容量(cap）两个字段。
	结构体作为函数参数时是值拷贝，同理，实际上slice作为函数参数时也是值拷贝，
在函数中对slice的修改是通过slice中保存的地址对底层数组进行修改，所以函数外的silce看起来被改变了。
当需要对slice做插入和删除时，由于需要更改长度字段，值拷贝就不行了，需要传slice本身在内存中的地址。
*/
