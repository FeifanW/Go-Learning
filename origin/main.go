package main

func main() {
	//// map的声明和注意事项
	//var a map[string]string
	//// 组使用map前，需要先make，make的作用就是给map分配数据空间
	//a = make(map[string]string, 10)
	//a["1"] = "宋江"
	//a["2"] = "吴用"
	//a["3"] = "武松"
	//a["4"] = "公孙胜"
	//fmt.Println(a)
	//
	//// 第二种方式
	//cities := make(map[string]string)
	//cities["1"] = "北京"
	//cities["2"] = "天津"
	//cities["3"] = "上海"
	//fmt.Println(cities)
	//cities["3"] = "上海~"
	//
	//// for-range遍历
	//for k, v := range cities {
	//	fmt.Printf("k=%v v=%v", k, v)
	//}
	//
	//fmt.Println(cities)
	//// 第三种方式
	//heroes := map[string]string{
	//	"hero1": "宋江",
	//	"hero2": "卢俊义",
	//}
	//fmt.Println(heroes)
	//
	///*
	//	练习：演示一个key-value 的value是map的案例
	//	比如：我们要存放3个学生信息，每个学生有name和sex信息
	//	思路：map[string]map[string]string
	//*/
	//studentMap := make(map[string]map[string]string)
	//
	//studentMap["stu01"] = make(map[string]string, 3)
	//studentMap["stu01"]["name"] = "tom"
	//studentMap["stu01"]["sex"] = "男"
	//studentMap["stu01"]["address"] = "北京长安街"
	//
	//studentMap["stu02"] = make(map[string]string, 3) // 不能少
	//studentMap["stu02"]["name"] = "mary"
	//studentMap["stu02"]["sex"] = "女"
	//studentMap["stu02"]["address"] = "上海黄浦江"
	//
	//fmt.Println(studentMap)
	//fmt.Println(studentMap["stu02"])
	//fmt.Println(studentMap["stu02"]["address"])

	// 演示map切片的使用
	/*
		要求：使用一个map来记录monster的信息 name 和 age，也就是说一个monster对应一个map，并且妖怪的个数可以动态的增加=>map切片
	*/
	// 1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	// 2.增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔"
		monsters[1]["age"] = "400"
	}
	// 这种写法越界
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string,2)
	//	monsters[2]["name"] = "狐狸"
	//	monsters[2]["age"] = "300"
	//}
	//这里我们需要使用到切片的append函数，可以动态的增加monster
	//1.先定义个monster信息
	//newMonster := map[string]string{
	//	"name": "新的妖怪",
	//	"age":  "200",
	//}
	//monsters = append(monsters, newMonster)
	//fmt.Println(monsters)

	////map的排序
	//map1 := make(map[int]int, 10)
	//map1[10] = 100
	//map1[1] = 13
	//map1[4] = 56
	//map1[8] = 90
	//fmt.Println(map1)
	//// 如果按照map的key的顺序进行排序输出
	////1.先将map的key放入到切片中
	////2.对切片排序
	////3.遍历切片，然后按照key来输出map的值
	//var keys []int
	//for k, _ := range map1 {
	//	keys = append(keys, k)
	//}
	//// 排序
	//sort.Ints(keys)
	//fmt.Println(keys)
	//
	//for _, k := range keys {
	//	fmt.Printf("map1[%v]=%v \n", k, map1[k])
	//}
}
