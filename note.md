#### 1、搭建Go开发环境

1. ##### 下载SDK

   下载地址：https://golang.google.cn/dl/

   - darwin是mac下的sdk
     - pkg是图形化安装包
     - tar.gz是解压就可以使用
   - freebsd是unix下的sdk
   - linux如果是32位系统：386.tar.gz 如果是64位系统，选择amd.tar.g
   - src.tar.gz是源码包
   - windows   .msi双击安装 .zip解压安装 32位选择 -386.zip   64位选择amd64.zip

   安装路径不要中文

   解压后bin里面是go的指令 go /  gofmt

   src里面是go的源码

2. ##### 配置环境变量

   ###### Windows下搭建Go开发环境：

   | 环境变量 | 说明                                 |
   | -------- | ------------------------------------ |
   | GOROOT   | 指定SDK的安装路径 d:/programs/go     |
   | Path     | 添加SDK的/bin目录                    |
   | GOPATH   | 工作目录，将来我们的go项目的工作路径 |

   注意：

   - Path这个环境变量不需要再创建，因为系统本身就有，你后面增加即可
   - 增加Go的bin：path里面添加一项 %GOROOT%\bin
   - 添加GOPATH变量：变量值填项目的路径

   测试环境是否配置成功：

   - go version

   注意：配置环境变量后，需要重新打开一次终端

   ###### Linux下搭建Go开发环境：

   32位：gox.x.x.linux-386.tar.gz

   64位：go.x.x.x.linux-amd64.tar.gz

   - 安装路径不要有中文
   - SDK安装建议：linux放在/opt目录下
   - 安装时，解压即可，我们使用的是tar.gz

   安装步骤：

   - uname -a 可以查看linux系统信息
   - 将安装包传输到linux系统
   - 把包拷贝到opt目录下  cp  包名 /opt
   - 权限不够的话切换到root角色 su root
   - tar -zxvf 包名  解压
   - 执行 ./go version测试一下

   配置环境变量：

   1. 在/etc/profile文件下添加三条语句

      export GOROOT=/opt/go

      export PATH=$GOROOT/bin:$PATH

      export GOPATH=$HOME/goprojects/

   2. 提示：修改/etc/profile文件，需要root权限或者sudo名字

   3. 步骤：

      - 使用root的权限来编辑 vim /etc/profile文件
      - 如果需要生效的话，需要注销一下，再使用

#### 2、开发基本结构说明

1. ##### windows下开发步骤

   - 安装windows的vscode
   - go代码写到xxx.go文件中 [可能需要和设置环境变量时设置的工作目录一致]
   - 通过go build命令对go文件进行编译，生成.exe文件
   - 在dos命令下执行.exe文件就可以看到运行效果
   - 注意：通过**go run** 命令可以直接运行xxx.go程序 [类似执行一个脚本文件的形式]

   **基本内容：**

   - go文件的后缀是go

   - package main 

     表示hello.go 文件所在的包是main，在go中，每个文件都必须属于一个包

   - import "fmt"

     表示：引入一个包，包名fmt，引入该包后，就可以使用fmt包的函数，比如：fmt.Println

   - func main(){

     }
     
     func是一个关键字，表示一个函数
     
     mian是函数名，是一个主函数，即我们程序的入口
     
   - fmt.Println('hello')

     表示调用fmt包的函数Println输出”hello,world“

   - go build xxx.go

     生成可执行文件

   - .\xxx.exe 即可执行上面生成的可执行文件

   - 也可以用go run执行上面的xxx.go

2. ##### Linux和Mac下开发Go程序

   说明：Linux开发Go和windows开发基本是一样的，只是在运行可执行的程序时，是以 **./文件名方式**

   也可以使用**go run**执行

3. ##### Go执行流程分析

   go build 然后执行

   go run是把编译运行合到一起了

   两种流程的方式区别：

   - 如果现编译生成了可执行文件，那么我们可以将该可执行文件拷贝到没有go开发环境的机器上，仍然可以运行
   - 如果直接go run go源代码，那么如果要在另外一台机器上运行，那么也需要go的开发环境，否则无法运行
   - 在编译时，编译器会将程序运行以来的库文件包含在可执行文件中，所以，可执行文件变大了

   如果想对生成的exe重命名，可以输入go build -o myhello.exe hello.go

   **编译：**

   - 编译就是将其编译成机器可以识别的二进制码文件
   - 在该源文件目录下，通过go build 对 hello.go文件进行编译，可以指定生成的可执行文件名，在windows下必须是.exe后缀
   - 如果程序没有错误，没有任何提示，会在当前目录下会出现一个可执行文件（windows下是.exe Linux下是一个可执行文件），该文件是二进制码文件，也是可以执行的程序
   - 如果程序有错误，编译时，会在错误的哪行报错

4. ##### Go的基本语法要求和注意事项

   - 源文件以go为扩展名
   - 应用程序的执行入口是main()函数
   - 严格区分大小写
   - Go方法由一条条语句构成，每个语句后不需要分号（Go语言会在每行后自动加分号）
   - Go编译器是一行行进行编译的，因此一行就写一条语句，不能把多条语句写在同一个，否则报错
   - go语言定义的变量或import的包如果没有使用到，代码不能编译通过
   - 大括号都是成对出现的，缺一不可

#### 3、基本语法及规则

1. ##### Go语言的转义字符

   | 转义字符 | 功能                                   |
   | -------- | -------------------------------------- |
   | \t       | 一个制表位，实现对齐的功能             |
   | \n       | 换行符                                 |
   | \\\      | 一个\                                  |
   | \\"      | 一个"                                  |
   | \r       | 一个回车 fmt.Println("星期五\r星期六") |

   





























