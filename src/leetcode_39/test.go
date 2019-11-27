package  main

import "fmt"

func main() {
	test1()
	test2()
}

func test1() {
	a := []int{1,2,3,4,5}
	b := a  //等号赋值
	fmt.Printf("b切片的值%v，a地址%p,b地址%p\n", b, a, b)
	a[0] = 100
	fmt.Printf("等号赋值当改变源值，新值也会改变 %v\n", b)
}

func test2() {
	a := []int{1,2,3,4,5}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Printf("b切片的值%v，a地址%p,b地址%p\n", b, a, b)
	a[0] = 100
	fmt.Printf("copy赋值当改变源值，新值不会改变 %v\n", b)
}



