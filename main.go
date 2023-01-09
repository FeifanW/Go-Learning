package main

import "fmt"

func main() {
	// map的声明和注意事项
	var a map[string]string
	// 组使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["1"] = "宋江"
	a["2"] = "吴用"
	a["3"] = "武松"
	a["4"] = "公孙胜"
	fmt.Println(a)

	// 第二种方式
	cities := make(map[string]string)
	cities["1"] = "北京"
	cities["2"] = "天津"
	cities["3"] = "上海"
	fmt.Println(cities)
	cities["3"] = "上海~"
	fmt.Println(cities)
	// 第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
	}
	fmt.Println(heroes)

	/*
		练习：演示一个key-value 的value是map的案例
		比如：我们要存放3个学生信息，每个学生有name和sex信息
		思路：map[string]map[string]string
	*/
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 3) // 不能少
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["address"])

}
