package main

import (
	"errors"
	"fmt"
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

	//var slice []float64 = make([]float64, 5, 10)
	//slice[1] = 10
	//slice[3] = 20
	//fmt.Println(slice)
	//fmt.Println("slice的size=", len(slice))
	//fmt.Println("slice的cap=", cap(slice))
	//
	//// 定义一个切片，直接就指定具体数组，使用原理类似make的方式
	//var strSlice []string = []string{"tom", "jack", "mary"}
	//fmt.Println("strSlice=", strSlice)
	//fmt.Println("strSlice size=", len(strSlice))
	//fmt.Println("strSlice=", cap(strSlice))

	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	//for i := 0; i < len(slice); i++ {
	//	fmt.Printf("slice[%v]=%v", i, slice[i])
	//}
	//fmt.Println()
	////使用for-range方式遍历切片
	//for i, v := range slice {
	//	fmt.Printf("i=%v v=%v \n", i, v)
	//}

	//用append内置函数，可以对切片进行动态增加
	var slice3 []int = []int{100, 200, 300}
	// 通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3)

	// 通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice...)
	fmt.Println("slice3", slice3)

	// 切片的拷贝操作
	// 切片使用copy内置函数完成拷贝，举例说明
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)

	str := "hello world"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)
	// 细节,转成[]byte后，可以处理英文和数字，但是不能处理中文
	// 原因是 []byte 字节来处理，而一个汉字是3个字节，因此就会出现乱码
	// 解决方法 是将 string 转成 []rune 即可， 因为 []rune 是按照字符处理，兼容汉字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str=", str)
}
