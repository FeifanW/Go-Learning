package main

import "fmt"

func main() {
	//// 基本数据类型在内存布局
	//var i int = 10
	//// i的地址是什么，&i
	//fmt.Println("i的地址=", &i)
	//
	//// 下面的var ptr *int = &i
	//// ptr是一个指针变量
	//// ptr的类型 *int
	//// ptr 本身的值&i
	//var ptr *int = &i
	//fmt.Printf("ptr=%v\n", ptr)

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("num =", num)
}
