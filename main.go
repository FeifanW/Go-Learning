package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 案例1：字符串转布尔值
	var str string = "true"
	var b bool
	/*
		b, _ = strconv.ParseBool(str)
		说明
		1.strconv.ParseBool(str) 函数会返回两个值 （value bool, err error)
		2.因为我只想获取到 value bool, 不想获取 err 所以我使用_忽略
	*/

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	// 案例2：字符串转int
	var str2 string = "12345690"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64) // 返回的是int64，想要得到int32需要再转一下
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	// 案例3：字符串转float
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)
}
