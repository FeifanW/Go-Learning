package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//// 打开文件
	//// 概念说明：file的叫法
	//// 1.file叫file对象
	//// 2.file叫file指针
	//// 3.file叫file文件句柄
	//file, err := os.Open("d:/test.txt")
	//if err != nil {
	//	fmt.Println("open file err=", err)
	//}
	//// 输出一下文件，看看文件是什么，看出file就是一个指针*File
	//fmt.Println("file=%v", file)
	//// 关闭文件
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("close file err=", err)
	//}
	//
	//// 当函数退出时，要及时的关闭file
	//defer file.Close() // 记着关闭不然会有内存泄露
	//
	//// 创建一个 *Reader，是带缓冲的
	///*
	//	const (
	//	defaultBufsize = 4096   // 默认的缓冲区为4096  u
	//	)
	//*/
	//reader := bufio.NewReader(file)
	//// 循环的读取文件的内容
	//for {
	//	str, err := reader.ReadString('\n') // 读到一个换行就结束
	//	if err == io.EOF {                  // io.EOF表示文件的末尾
	//		break
	//	}
	//	// 输出内容
	//	fmt.Printf(str)
	//}
	//fmt.Println("文件读取结束")

	// 使用ioutil.ReadFile一次性将文件读取到位
	//file := "d:/test.txt"
	//content, err := ioutil.ReadFile(file)
	//if err != nil {
	//	fmt.Printf("read file err=%v", err)
	//}
	// 把读取到的内容显示到终端
	//fmt.Printf("%v", content)         //[]byte   输出的都是切片
	//fmt.Printf("%v", string(content)) //[]byte   输出的都是数组
	// 因为，我们没有显示Open文件，因此也不需要显示的Close文件
	// 因为，文件的Open和Close被封装到ReadFile 函数内部

	// 创建一个新文件，写入内容 5句 "hello,Gardon"
	// 1.打开文件 d:/abc.txt
	filePath := "d:/abc.txt"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) // 第二个表示清空文件内容
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) // 第二个表示清空文件内容
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666) // 第二个表示清空文件内容
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 及时关闭file句柄
	defer file.Close()
	// 先读取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 如果读取到文件的末尾
			break
		}
		// 显示到终端
		fmt.Print(str)
	}
	// 准备写入5句"hello Gardon"
	str := "TODAY\r\n"
	// 写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	// 因为Writer是带缓存，因此在调用WriterString方法时，其实
	// 内容是先写入到缓存的，所以需要调用Flush方法，将缓冲的数据
	// 真正的写入到文件中，否则文件中会没有数据
	writer.Flush()
}
