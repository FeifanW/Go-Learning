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

type Person struct {
	Name string
}

// 给Person类型绑定一方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func main() {

	var p Person
	p.Name = "tom"
	p.test()

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
