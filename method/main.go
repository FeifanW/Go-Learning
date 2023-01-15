package main

import "fmt"

//type Point struct {
//	x int
//	y int
//}
//
//type Rect struct {
//	leftUp, rightDown Point
//}
//type Rect2 struct {
//	leftUp, rightDown *Point
//}
//
//type A struct {
//	Num int
//}

//type B struct {
//	Num int
//}
//type Monster struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Skill string `json:"skill"`
//}

//type Person struct {
//	Name string
//}

// 给Person类型绑定一方法
//func (p Person) test() {
//	fmt.Println("test() name=", p.Name)
//}

//func (p Person) test03() {
//	p.Name = "jack"
//	fmt.Println("test03() =", p.Name)
//}
//func (p *Person) test04() {
//	p.Name = "mary"
//	fmt.Println("test04() =", p.Name)
//}

type Stu struct {
	Name string
	Age  int
}

func main() {

	// 方式1
	// 在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}
	//在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu3 = Stu{
		Name: "jack",
		Age:  20,
	}
	stu4 := Stu{
		Age:  30,
		Name: "mary",
	}
	fmt.Println(stu1, stu2, stu3, stu4)

	//方法2， 返回结构体的指针类型
	var stu5 = &Stu{"小王", 29}
	stu6 := &Stu{"小王~", 39}
	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法不依赖字段的定义顺序
	var stu7 = &Stu{
		Name: "小李",
		Age:  49,
	}
	stu8 := &Stu{
		Age:  59,
		Name: "小李~",
	}
	fmt.Println(*stu5, *stu6, *stu7, *stu8)

	//p := Person{"tom"}
	//p.test03()
	//fmt.Println("main() p.name=", p.Name) //tom
	//(&p).test03()                         // 从形式上传入地址，但本质仍然是值拷贝
	//fmt.Println("main() p.name=", p.Name) // tom
	//(&p).test04()
	//fmt.Println("main() p.name=", p.Name) // mary
	//p.test04()                            // 等价于(&p).test04()  编译器自动处理，从形式上是传入值类型，但是本质任然是地址拷贝

	//var p Person
	//p.Name = "tom"
	//p.test()

	////1.创建一个Monster变量
	//monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	////2.将monster变量序列化为json格式字串
	//// json.Marshal 函数中使用反射
	//jsonStr, err := json.Marshal(monster)
	//if err != nil {
	//	fmt.Println("json处理错误", err)
	//}
	//fmt.Println("jsonStr", string(jsonStr))
	//
	//var a A
	//var b B
	//a = A(b) // 可以转换，但是有要求，就是结构体的字段要完全一样（包括：名字、个数和类型）
	//fmt.Println(a, b)
	//
	//r1 := Rect{Point{1, 2}, Point{3, 4}}
	//// r1有四个int，在内存中是连续分布
	//// 打印地址
	//fmt.Printf("r1.leftUp.x 地址是=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p  \n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	//
	//// r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的
	//// 但是他们指向的地址不一定是连续的
	//r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	//// 打印地址
	//fmt.Printf("r2.leftUp 本身地址是=%p r2.rightDown 本身地址=%p  \n", &r2.leftUp, &r2.rightDown)
	//fmt.Printf("r2.leftUp 指向地址是=%p r2.rightDown 指向地址=%p  \n", r2.leftUp, r2.rightDown)

}
