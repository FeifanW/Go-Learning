package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	// 1.统计字符串的长度，按字节len(str)
	str := "hello北" // golang的编码统一为utf-8 (ascii的字符（字母和数字）)上一个字节  汉字
	fmt.Println("str len=", len(str))

	// 2.字符串遍历，同时处理有中文的问题 r := []rune(str)
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符串=%c\n", r[i])
	}

	// 3.字符串转整数：n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	// 4.整数转字符串 str = strconv.Itoa(12345)
	str3 := strconv.Itoa(123456)
	fmt.Printf("str=%v, str=%T", str3, str3)

	// 5.字符串 转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// 6.[]byte 转 字符串 str = string([]byte{97,98,99})
	var str4 = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str4)

	// 7.10进制转2,8,16进制: str = strconv.FormatInt(123,2) 返回对应的字符串
	var str5 = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str5)

	// 8.查找子串是否在指定进制的字符串中 :strings.Contains("seafood","foo")
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	// 9.统计一个字符串有几个指定的子串  :strings.Count("ceheese","e")
	num := strings.Count("ceheese", "e")
	fmt.Printf("b=%v\n", num)

	// 10.不区分大小写的字符串比较（==是区分字母大小写的）:fmt.Println(strings.EqualFold("abc","Abc"))
	fmt.Println(strings.EqualFold("abc", "Abc"))
	fmt.Println("结果", "abc" == "Abc") // 区分字母大小写

	// 11.返回子串在字符串中第一次出现的index值，如果没有返回-1 :strings.Index("NLT_abc","abc")
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v\n", index)

	// 12.返回子串在字符串最后一次出现的index，如果没有返回-1：strings.LastIndex("go golang","go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	// 13.将指定的子串替换成另一个子串：strings.Replace("go go hello","go","go语言",n) n 可以指定你希望替换几个，如果n=-1表示全部替换
	str = strings.Replace("go go hello", "go", "go语言", 1) // 第一个参数可以传一个变量
	fmt.Printf("str=%v\n", str)

	// 14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world",",")
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	// 15.将字符串的字母进行大小写的转换：strings.ToLower("Go")  // go strings.ToUpper("Go")
	str = "goLang hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	// 16.将字符串左右两边的空格去掉：strings.TrimSpace("  javascript     ")
	str = strings.TrimSpace("  javascript     ")
	fmt.Printf("str=%v\n", str)

	// 17.将字符串左右两边指定的字符去掉，只想去左边就用TrimLeft  去右边就用TrimRight
	str = strings.Trim("! hello !", " !")
	fmt.Printf("str=%q\n", str)

	// 18.判断字符串是否以指定的字符串开头 strings.HasPrefix("ftp://192.168.10.1","ftp")
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
}
