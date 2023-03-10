package main

import (
	"Go-Learning/customerManage/model"
	"Go-Learning/customerManage/service"
	"fmt"
)

type customerView struct {
	// 定义必要的字段
	key  string // 接收用户输入
	loop bool   // 表示是否循环的显示主菜单
	// 增加一个字段customerService
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {
	// 首先，获取到当前所有的客户信息（在切片中）
	customers := this.customerService.List()
	// 显示
	fmt.Println("-----------------客户列表------------------")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n----------------客户列表完成----------------")
}

// 得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("-------------------添加客户--------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的Customer实例
	// 注意：id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	// 调用
	if this.customerService.Add(customer) {
		fmt.Println("--------------添加完成----------------")
	} else {
		fmt.Println("--------------添加失败-----------------")
	}
}

// 得到用户的输入id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("------------------删除客户--------------------")
	fmt.Println("请选择删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除操作
	}
	// 这里可以加入一个循环判断，直到用户输入y或者n才退出
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		// 调用customerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------------删除完成----------------------")
		} else {
			fmt.Println("------------------删除失败，输入的id号不存在--------------------")
		}
	}

}

// 退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("------------------客户信息管理软件------------------------")
		fmt.Println("                    1.添加客户")
		fmt.Println("                    2.修改客户")
		fmt.Println("                    3.删除客户")
		fmt.Println("                    4.客户列表")
		fmt.Println("                    5.退   出")
		fmt.Println("请选择（1-5）")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	// 在main函数中，创建一个customerView,并运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	// 这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	// 显示主菜单
	customerView.mainMenu()
}
