package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其他字符的个数
}

// 自己编写一个函数，接收两个文件路径 srcFileName  dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	// 通过srcfile，获取到reader
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 通过dstFile,获取到Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main() {
	/*
			// 将d:/abc.txt 文件内容导入到 d:/test.txt中
			// 1.首先将 d:/abc.txt内容读取到内存
			// 2.将读取到的内容写入 d:/test/txt
			file1Path := "d:/abc.txt"
			file2Path := "d:/kkk.txt"

			data, err := ioutil.ReadFile(file1Path)
			if err != nil {
				// 说明读取文件有错误
				fmt.Printf("read file err=%v", err)
				return
			}
			err = ioutil.WriteFile(file2Path, data, 0666)
			if err != nil {
				fmt.Printf("write file error=%v\n", err)
			}




		srcFile := "d:/abc.txt"
		dstFile := "d:/ccc.txt"
		_, err := CopyFile(dstFile, srcFile)
		if err == nil {
			fmt.Println("拷贝完成")
		} else {
			fmt.Println("拷贝错误 err=%v", err)
		}


	*/

	// 思路：打开一个文件，创一个Reader
	// 每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	// 然后将结果保存到一个结构体
	fileName := "d:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close() // 打开之后就要及时关闭
	// 定义个charCount实例
	var count CharCount
	// 创建一个Reader
	reader := bufio.NewReader(file)
	// 开始循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 读到文件末尾就退出
			break
		}
		// 遍历str,进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v < '9':
				count.NumCount++
			default:
				count.OtherCount++
			}

		}
	}
	// 输出统计的结果看看
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
