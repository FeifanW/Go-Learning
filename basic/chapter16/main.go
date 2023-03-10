package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name string
	//Age      int
	//Birthday string
	//sal      float64
	//skill    string
}

// 演示将json字符串，反序列化成struct
func unmarshalStruct() {
	// 说明str 在项目开发中，是通过网络传输获取到
	str := "{\"name\":\"jack\"}"
	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n", monster)
}

// 将json字符串反序列化成map
func unmarshalMap() {
	str := "{\"name\":\"jack\"}" // 如果是程序读取的，则不用加\号
	// 定义一个map
	var a map[string]interface{}

	// 反序列化
	// 反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

// 演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"name\":\"jack\"}]"
	// 定义一个slice
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
