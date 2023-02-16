package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
		fmt.Println("命令行的参数有", len(os.Args))
		// 遍历os.Args切片，就可以得到所有的命令行输入参数值
		for i, v := range os.Args {
			fmt.Printf("args[%v]=%v\n", i, v)
		}
	*/
	// 定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	// &user就是接收用户命令行输入的-u后面的参数值
	// "u" 就是-u指定参数
	// "" 默认值
	// "用户名，默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	// 这里有一个非常重要的操作，转换，必须调用该方法
	flag.Parse()
	// 输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
