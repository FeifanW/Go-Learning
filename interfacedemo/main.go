package main

import "fmt"

//// 声明/定义一个接口
//type Usb interface {
//	// 声明了两个没有实现的方法
//	Start()
//	Stop()
//}
//type Phone struct {
//}

// 让Phone实现Usb接口的方法
//func (p Phone) Start() {
//	fmt.Println("手机开始工作")
//}
//func (p Phone) Stop() {
//	fmt.Println("手机停止工作")
//}
//
//type Camera struct {
//}
//
//// 让Camera实现 Usb接口的方法
//func (c Camera) Start() {
//	fmt.Println("相机开始工作")
//}
//func (c Camera) Stop() {
//	fmt.Println("相机停止工作")
//}
//
//// 计算机
//type Computer struct {
//}
//
//// 编写一个方法Working方法，接收一个Usb接口类型变量
//// 只要实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明的所有方法）
//func (c Computer) Working(usb Usb) {
//	// 通过usb接口变量来调用Start和Stop方法
//	usb.Start()
//	usb.Stop()
//}

type Usb interface {
	Say()
}
type Stu struct {
}

func (this *Stu) Say() {
	fmt.Println("Say()")
}
func main() {
	var stu Usb = &Stu{}
	//var stu Stu = Stu{}
	// 错误！会报Stu类型没有实现Usb接口
	// 如果希望通过编译 var u Usb = &stu
	var u Usb = stu
	u.Say()
	fmt.Println("here", u)

	////测试
	////先创建结构体变量
	//computer := Computer{}
	//phone := Phone{}
	//camera := Camera{}
	////关键点
	//computer.Working(phone)
	//computer.Working(camera)
}
