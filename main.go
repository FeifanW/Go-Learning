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

	var name string
	var age byte
	var sal float32
	var isPass bool

	// 方法1
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
	//// 方法2
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}
