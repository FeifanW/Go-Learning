package main

import (
	"fmt"
	"strconv"
	"time"
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

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行时间为%v\n", end-start)

}
