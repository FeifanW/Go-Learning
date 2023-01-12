package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
}

//定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func modifyUser(users map[string]map[string]string, name string) {
	// 判断users中是否有name
	if users[name] != nil {
		// 有这个用户
		users[name]["pws"] = "888888"
	} else {
		// 没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name

	}

}

func main() {
	// map是引用类型，遵守引用类型传递的机制，在一个函数接收map
	// 修改后，会直接修改原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)

	//map的value 也经常使用struct类型
	//更适合管理复杂的数据（比前面的value是一个map更好）
	//比如value为Student结构体
	//1.map的key为学生的学号，是唯一的
	//2.map的value为结构体，包含学生的姓名，年龄，地址
	students := make(map[string]Stu, 10)
	// 创建2个学生
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 20, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)
	// 遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
		fmt.Println()
	}

	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println(users)
}
