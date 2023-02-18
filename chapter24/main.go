package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改
// num int的值
// 修改student的值
func reflect01(b interface{}) {
	// 获取到reflect.value
	rVal := reflect.ValueOf(b)
	// 看看rVal的Kind是
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	// rVal
	rVal.Elem().SetInt(20) // 会修改rVal指针指向的那个值
}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)

	// 可以这样理解rVal.Elem()
	//num := 9
	//ptr *int = &num
	//num2 := *ptr // ===类似 rVal.Elem()
}
