package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程（可以理解成进程）中，开启一个goroutine，该协程每隔1秒输出"hello world"
// 在主线程中也每隔一秒输出"hello,golang",输出10次后退出程序
// 要求主线程和goroutine同时执行

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world" + strconv.Itoa(i)) // strconv.Itoa() 将数字转成字符串
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
