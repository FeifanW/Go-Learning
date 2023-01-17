package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()

	// 嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Name:  "电视机",
			Price: 5000.99,
		},
		Brand{
			Name:    "海尔",
			Address: "山东",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{&Goods{"电视机", 7000.5}, &Brand{"创维", "河南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视机",
			Price: 7000.5,
		},
		&Brand{"创维", "河南"}}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
}
