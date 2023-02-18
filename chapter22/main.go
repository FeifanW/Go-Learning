package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	// 关闭intChan
	close(intChan)
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	// 这里可以使用defer + recover
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	// 定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

// 从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for循环
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { // intChan取不到
			break
		}
		flag = true // 假设是素数
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数放入到primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	// 这里还不能关闭primeChan
	// 向exitChan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	// 标识退出的管道
	exitChan := make(chan bool, 4) // 4个
	// 开启一个协程，向intChan放入 1-8000 个数
	go putNum(intChan)
	// 开启4个协程，从intChan取出数据，并判断是否为素数，如果是就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	// 这里主线程，进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// 当我们从exitChan 取出4个结果，就可以放心的关闭primeNum
		close(primeChan)
	}()
	// 比那里我们的primeNum，把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}
