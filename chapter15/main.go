package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	sal      float64
	skill    string
}

func testStruct() {
	// 演示
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		sal:      8000.0,
		skill:    "牛魔拳",
	}
	// 将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

// 将map序列化
func testMap() {
	// 定义一个map
	var a map[string]interface{}
	// 使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "孙悟空"
	a["age"] = 25
	a["address"] = "水帘洞"
	// 将a这个map进行序列化
	// 将monster序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("a map序列化后=%v\n", string(data))
}

// 演示对切片进行序列化，我们这个切片 []map[string]interface{}
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	// 使用map前，需要先make
	m2 = make(map[string]interface{})
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["address"] = [2]string{"墨西哥", "阿根廷"}
	slice = append(slice, m2)
	// 将切片序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v", err)
	}
	// 输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 对基本数据类型序列化，没有什么实际意义
func testFloat64() {
	var num1 float64 = 2345.67
	// 对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("基本数据类型 序列化后=%v\n", string(data))
}

func main() {
	// 演示将结构体，map,切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
