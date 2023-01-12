package main

import "fmt"

//定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

// 如果结构体的字段类型是：指针、slice和map的零值都是nil，即还没有分配空间
// 如果需要使用这样的字段，需要先make,才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string // 切片
}
type Monster struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string // 切片
}

func main() {
	// 创建一个Cat变量
	//var cat1 Cat
	//cat1.Name = "小白"
	//cat1.Age = 3
	//cat1.Color = "白色"
	//cat1.Hobby = "吃鱼"
	//fmt.Println("cat1=", cat1)
	//fmt.Println("猫猫的信息如下", "")
	//fmt.Println("Age=", cat1.Age)
	//fmt.Println("color=", cat1.Color)
	//fmt.Println("hobby=", cat1.Hobby)
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	// 使用slice 再次说明，一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// 使用map,一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"
	fmt.Println(p1)

	//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改不影响另外一个，结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 // 结构体是值类型，默认为值拷贝
	//monster2 := &monster1 // 如果想改同一个，则传地址
	monster2.Name = "青牛精"

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
}
