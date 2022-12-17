package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	// 直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 如果我们希望输出对应的字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	// var c3 byte = "北"   // overflow溢出
	var c3 int = '北' // overflow溢出
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3, c3)

	var a int
	var b float32
	var c float64
	var isMarried bool
	var name string
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v", a, b, c, isMarried, name)
}
