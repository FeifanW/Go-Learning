package main

import "fmt"

func main() {
	/*
		// 演示一下管道的使用
		// 1.创建一个可以存放3个int类型的管道
		var intChan chan int // 本身是一个地址
		intChan = make(chan int, 3)
		// 2.看看intChan是什么
		fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)
		// 3.向管道写入数据
		intChan <- 10
		num := 211
		intChan <- num
		// 注意，当我们给管道写入数据时，不能超过其容量
		// 4.看看管道的长度和cap(容量)
		fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
		// 5.从管道中读取数据
		var num2 int
		num2 = <-intChan
		fmt.Println("num2=", num2)
		fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))
		// 6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告deadlock
		num3 := <-intChan
		num4 := <-intChan
		fmt.Println("num3=", num3, "num4=", num4)
	*/
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // 关闭管道
	// 不能够再写入了 intChan <- 300
	fmt.Println("okok")
	// 当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入100个数据到管道
	}

	// 在遍历时，如果channel没有关闭，则会出现deadlock的错误
	close(intChan2)
	// 遍历管道不能使用for循环
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

}
