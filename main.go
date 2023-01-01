package main

import (
	"errors"
	"strconv"
)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

/*

 */
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取...
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		// 如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
}

func main() {

}
