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

#### 3、Go基础内容

1. ##### Go语言的转义字符

   | 转义字符 | 功能                                                         |
   | -------- | ------------------------------------------------------------ |
   | \t       | 一个制表位，实现对齐的功能                                   |
   | \n       | 换行符                                                       |
   | \\\      | 一个\                                                        |
   | \\"      | 一个"                                                        |
   | \r       | 一个回车 fmt.Println("星期五\r星期六")，从当前行的最前面开始输出，覆盖掉以前内容 |

2. ##### Go开发常见错误和解决方法

   - 找不到文件：源文件名不存在或者写错，或者当前路径错误
   - 语法错误：编译器会报告错误信息

3. ##### 注释

   - 行注释  //
   - 块注释 /* */

   注意：块注释里面不允许再有块注释嵌套

4. ##### 规范的代码风格

   - 推荐使用行注释注释整个方法和语句

   - 运算符两边习惯性加一个空格

   - 花括号要这样写

     ```
     func main(){
     
     }
     ```

   - 一行不超过80个字符，超过请使用换行展示

5. ##### Dos常用指令

   **目录操作：**

   Disk Operating System 磁盘操作系统

   终端指令  ----------->  中间的Dos系统解析指令  ----------->  文件操作

   - 查看当前目录 dir

   - 切换到其他盘，比如F盘    cd /d f:

   - 切换到当前盘的其他目录下（相对路径和绝对路径演示）

     绝对路径：从当前盘符最上面开始定位，比如D:F:C:找到对应的目录，比如d:\test

     相对路径：从当前位置开始定位，找对应的目录

   - 切换到上一级

     ..代表上级目录  cd ..

   - 切换到根目录

     cd \

   - 新建目录md

     md test1 test2      这是新建多个文件夹

   - 删除目录rd

     rd test1     如果这个文件夹里面有内容，无法删除成功

     rd /q/s test1  这里的q代表不用询问，s代表下面的所有目录

     rd /s test1   带询问的删除

   **文件操作：**

   - 新建或追加内容到文件

     echo hello > d:\test\abc.txt

   - 复制或移动文件

     copy abc.txt d:\test   拷贝时使用原来的名字

     copy abc.txt d:\test\ok.txt   拷贝时重新指定名称

     move abc.txt f:\    移动文件也可以理解为剪切

   - 删除文件

     del abc.txt

     del *.txt    删除所有txt后缀的文件

   ##### 其他指令：

   - 清屏 cls
   - 退出 dos  exit

6. ##### 变量

   内存中一个数据存储空间的表示

   Go变量使用的三种方式：

   - 第一种：指定变量类型，声明后若不赋值，使用默认值

   - 第二种：根据值自行判断变量类型（类型推导）

   - 第三种：省略var 注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误

     ```go
     // 等价于var name string  name = "tom"
     name := "tom"
     ```

   多变量声明

   ```go
   // 方式1
   var n1, n2, n3 int
   // 方式2
   var n1, name, n3 = 100, "tom", 888
   ```

   定义全局变量

   ```go
   // 方式1
   var n1 = 100
   var n2 = 200
   var name = "jack"
   // 方式2
   var(
   	n3 = 300
       n4 = 900
       name2 = "mary"
   )
   ```

   注意事项：

   - 该区域的数据值可以在同一类型范围内不断变化
   - 变量在同一个作用域内不能重名
   - 变量 = 变量名 + 值 + 数据类型，这一点大家注意
   - Golang的变量如果没有赋初值，编译器会使用默认值，比如int默认值0  string默认值为空串

   \+ 号的使用：

   - 当左右都是数值型时，则做加法运算
   - 当左右两边都是字符串时，则做字符串拼接

7. ##### 数据类型

   ![image-20221216143331391](D:\practice Space\Go-Learning\assets\image-20221216143331391.png)

###### 整数的类型：

有符号：

| 类型  | 有无符号 | 占用存储空间 | 表数范围     |
| ----- | -------- | ------------ | ------------ |
| int8  | 有       | 1字节        | -128 ~ 127   |
| int16 | 有       | 2字节        | -2^15~2^15-1 |
| int32 | 有       | 4字节        | -2^31~2^31-1 |
| int64 | 有       | 8字节        | -2^63-2^63-1 |

无符号：

| 类型   | 有无符号 | 占用存储空间 | 表数范围 |
| ------ | -------- | ------------ | -------- |
| uint8  | 无       | 1字节        | 0~255    |
| uint16 | 无       | 2字节        | 0~2^16-1 |
| uint32 | 无       | 4字节        | 0~2^32-1 |
| uint64 | 无       | 8字节        | 0-2^64-1 |

其他：

| 类型 | 有无符号 | 占用存储空间                     | 表数范围                        | 备注                         |
| ---- | -------- | -------------------------------- | ------------------------------- | ---------------------------- |
| int  | 有       | 32位系统4个字节，64位系统8个字节 | -2^31~2^31-1  <br/>-2^63~2^63-1 |                              |
| uint | 无       | 32位系统4个字节，64位系统8个字节 | 0~2^32-1  <br/>0~2^64-1         |                              |
| rune | 有       | 与int32一样                      | -2^31~2^31-1                    | 等价int32，表示一个Unicode码 |
| byte | 无       | 与uint8等价                      | 0~255                           | 当要存储字符时选用byte       |

如何在程序查看某个变量的占用字节大小和数据类型（使用较多）

```go
package main
import (
	"fmt",
    "unsafe"
)
var n2 int64 = 10
// unsafe.SizeOf(n2) 是unsafe包的一个函数，可以返回n1变量占用的字节数
fmt.Printf("n2 的类型 %T n2占用的字节数是 %d", n2, unsafe.Sizeof(n2))
```

Go程序中整形变量在使用时，遵循保小不保大的原则，保证程序正确运行下，尽量使用占用空间小的数据类型

bit是计算机中最小存储单位，byte计算机中的基本存储单元  1btye = 8bit

Go的整型默认声明为int型

###### 小数类型：

| 类型          | 占用存储空间 | 表数范围             |
| ------------- | ------------ | -------------------- |
| 单精度float32 | 4字节        | -3.403E38~3.403E38   |
| 单精度float64 | 8字节        | -1.798E308~1.798E308 |

说明：

1. 浮点数在机器中存放形式简单说明：浮点数=符号位+指数位+尾数位

2. 尾数部分可能丢失，造成精度损失 -123.0000901

   float64的精度比float32要准确

3. 浮点数的存储分为三部分“符号位+指数位+尾数位  在存储过程中，精度会有丢失

4. 在相同的占用字节大小的情况下，浮点数比整数能够存储的数据更大

5. 浮点数都是有符号的

细节：

1. Go浮点类型有固定的范围和字段长度，不受具体OS（操作系统）的影响
2. Go的浮点类型默认声明为float64类型
3. 浮点型常量有两种表示形式
   - 十进制数形式：如5.12  .512（必须有小数点）
   - 科学计数法形式 5.123e2 = 5.12*10的2次方  5.12E-2 = 5.12/10的2次方
4. 通常情况下，应该使用float64，因为它比float32更精确，开发中推荐float64

###### 字符类型：

Go中没有专门的字符类型，如果存储单个字符（字母），一般使用byte来保存

传统的字符串是由字符组成，而**Go的字符串不同，是由字节组成**

```go
var c1 byte = 'a'
var c2 byte = '0'
// 直接输出byte值，就是输出了对应的字符的码值
fmt.Println("c1=",c1)
fmt.Println("c2=",c2)
// 如果我们希望输出对应的字符，需要使用格式化输出
fmt.Printf("c1=%c c2=%c\n",c1, c2)
// var c3 byte = "北"   // overflow溢出
var c3 int = '北'    // overflow溢出
fmt.Printf("c3=%c c3对应的码值=%d",c3, c3)
```

说明：

- 如果我们保存的字符在ASCII表的，比如[0-1,a-z,A-Z...]直接可以保存到byte
- 如果我们保存的字符对应码值大于255，这时我们可以考虑使用int类型保存
- 如果我们需要安装字符的方式输出，这时我们需要格式化输出，即fmt.Printf("%c",c1)

细节：

1. 字符常量是单引号('')括起来的单个字符，例如var c1 byte = 'a'  var c2 int = '中'
2. Go中允许使用转义字符'\\'来将其后的字符串转变为特殊字符型常量  var c3 char = '\\n'
3. Go语言的字符使用的是UTF-8编码
4. 在Go中，字符本质是一个整数，直接输出时，是该字符对应的UTF-8编码的码值
5. 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode字符
6. 字符类型是可以进行运算的，相当于一个整数，因为它都对应有Unicode码

字符串本质探讨：

1. 字符型存储到计算机中，需要将字符对应的码值（整数）找出来

   存储：字符 --->  对应码值  ---> 二进制  ----> 存储

   读取：二进制 ---> 码值 ---> 字符 ---> 读取

2. 字符和码值的对应关系是通过字符编码表决定的（是规定好）

3. Go语言的编码都统一成了utf-8，非常方便，很统一，再也没有编码乱码的困扰了

###### 布尔类型：

基本：

- 布尔类型也叫bool类型，bool类型数据只允许取值true和false
- bool类型占1个字节
- bool类型适用于逻辑运算，一般用于程序流程控制
  - if条件控制语句
  - for循环控制语句

###### 字符串类型：

字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的，Go语言的字符串的字节使用UTF-8编码标识Unicode文本

注意：

1. Go统一使用UTF-8编码

2. 字符串一旦赋值了，字符串就不能修改了，**在Go中字符串是不可改变的**

3. 字符串的两种表示形式

   - 双引号，会识别转义字符
   - 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果

4. 字符串拼接方式 + 号

   当一个拼接操作很长时，可以分行写，但是要注意需要将+保留在上一行

###### 基本数据类型默认值：

| 数据类型 | 默认值 |
| -------- | ------ |
| 整型     | 0      |
| 浮点型   | 0      |
| 字符串   | ""     |
| 布尔类型 | false  |

```go
var a int 
var b float32
var c float64
var isMarried bool
var name string
fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v",a,b,c,isMarried,name)
```

###### 基本数据类型的转换：

Go和java/c不同，Go在不同类型的变量之间赋值时需要**显示转换**，也就是说Golang中的数据类型**不能自动转换**

基本语法：

表达式T(v)将值v转换为类型T

T：就是数据类型，比如int32，int64，float32等等

v：就是需要转换的变量

```go
var i int32 = 100
// 希望将 i => float
var n1 float32 = float32(i)
var n2 int8 = int8(i)
var n3 int64 = int64(i) // 低精度 -> 高精度
fmt.Printf("i=%v n1=%v n2=%v n3=%v", i, n1, n2, n3)
fmt.Printf("i type is %T\n",i)    // int32
```

细节：

1. Go中，数据类型的转换可以是从表示范围小 -> 表示范围大，也可以 范围的 -> 范围小
2. 被转换的是**变量存储的数据**（即值），原来那个变量本身的数据类型并没有变化！
3. 在转换中，比如将int64转成int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样

如果没有使用到一个包，但是有想去掉，前面加一个_

```go
import (
	_ "fmt"
)
```

```go
var n1 int32 = 12
var n3 int8
var n4 int8
n4 = int8(n1) + 127  // 编译通过，但是结果不是127 + 12 按溢出处理
n3 = int8(n1) + 128  // 编译不过
fmt.Println(n3)
```

###### 基本数据类型转string:

- 方法1：fmt.Sprintf("%参数",表达式)

  1. 参数需要和表达式数据类型相匹配
  2. fmt.Sprintf() 会返回转换后的字符串

  ```go
  var num1 int = 99
  var num2 float64 = 23.456
  var b bool = true
  var myChar byte = 'h'
  var str string // 空的str
  
  str = fmt.Sprintf("%d", num1)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%f", num2)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%t", b)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = fmt.Sprintf("%d", myChar)
  fmt.Printf("str type %T str=%q\n", str, str)
  // 具体关于%号加字母代表的含义，去查Go语言官网文档fmt函数
  ```

- 方法2：使用strconv包的函数

  1. FormatBool
  2. FormatInt
  3. FormatUint
  4. FormatFloat
  
  ```go
  var num3 int = 99
  var num4 float64 = 23.456
  var b2 bool = true
  var str string // 空的str
  
  str = strconv.FormatInt(int64(num3), 10)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  // strconv.FormatFloat(num4, 'f', 10, 64)
  // 说明：'f'格式 10: 表示小数位保留10位   64:表示这个小数是float64
  str = strconv.FormatFloat(num4, 'f', 10, 64)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  str = strconv.FormatBool(b2)
  fmt.Printf("str type %T str=%q\n", str, str)
  
  // strconv包中有一个函数Itoa
  var num5 int64 = 4567
  str = strconv.Itoa(int(num5))
  fmt.Printf("str type %T str=%q\n", str, str)
  ```

###### string类型转基本数据类型：

1. 使用时strconv包的函数

   - ParseBool
   - ParseFloat
   - ParseInt
   - ParseUint

2. 案例：

   ```go
   // 案例1：字符串转布尔值
   var str string = "true"
   var b bool
   /*
   b, _ = strconv.ParseBool(str)
   说明
   1.strconv.ParseBool(str) 函数会返回两个值 （value bool, err error)
   2.因为我只想获取到 value bool, 不想获取 err 所以我使用_忽略
   */
   
   b, _ = strconv.ParseBool(str)
   fmt.Printf("b type %T b=%v\n", b, b)
   
   // 案例2：字符串转int
   var str2 string = "12345690"
   var n1 int64
   var n2 int
   n1, _ = strconv.ParseInt(str2, 10, 64) // 返回的是int64，想要得到int32需要再转一下
   n2 = int(n1)
   fmt.Printf("n1 type %T n1=%v\n", n1, n1)
   fmt.Printf("n2 type %T n2=%v\n", n2, n2)
   
   // 案例3：字符串转float
   var str3 string = "123.456"
   var f1 float64
   f1, _ = strconv.ParseFloat(str3, 64)
   fmt.Printf("f1 type %T f1=%v\n", f1, f1)
   ```

注意：

在将String类型转成基本数据类型时，要确保String类型能够转成有效的数据，比如我们可以把"123"转成一个整数，但是不能把"hello"转成一个整数，这样做Go会直接将其转成0

```go
var str4 string = "hello"
var n3 int64
n3, _ = strconv.ParseInt(str4, 10, 64)
fmt.Printf("n3 type %T n3=%v\n", n3, n3)
```

###### 指针：

1. 基本数据类型，变量存的就是值，也叫值类型
2. 获取变量的地址，用&，比如：var num int，获取num的地址：&num
3. 指针类型，变量存的是一个地址，这个地址指向的空间存的才是值 比如：var ptr *int = &num
4. 获取指针类型所指向的值，使用：*，比如：var ptr \*int，使用* \*ptr获取ptr指向的值
5. *是指针运算符，可以表示一个变量是指针类型，也可以表示一个指针变量所指向的存储单元，也就是这个地址所存储的值
6. &是取址符号，即取得某个变量的地址

```go
// 基本数据类型在内存布局
var i int = 10
// i的地址是什么，&i
fmt.Println("i的地址=", &i)

// 下面的var ptr *int = &i
// ptr是一个指针变量
// ptr的类型 *int
// ptr 本身的值&i
var ptr *int = &i
fmt.Printf("ptr=%v\n", ptr)
```

```go
var ptr *int
ptr = &num
*ptr = 10
fmt.Println("num =", num)
```

指针细节：

1. 值类型，都有对应的指针类型，形式为 ***数据类型**，比如 int的对应的指针就是 \*int，float32对应的指针类型就是 \*float，依次类推
2. 值类型包括：基本数据类型**int系列，float系列，bool，string，数组**和**结构体struct**

###### 常见的值类型和引用类型：

1. 值类型：基本数据类型int系列、float系列、bool、string、数组和结构体struct
2. 引用类型：指针、slice切片、map、管道chan、interface等都是引用类型

特点：

- 值类型：变量直接存储值，内存通常在栈中分配
- 引用类型：变量存储的是一个地址，这个地址对应的空间才是真正存储数据（值），内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾由GC回收

内存里面分为栈区和堆区：

栈区：值类型数据，通常是在栈区

堆区：引用类型，通常在堆区分配空间

###### 标识符的命名规范（重点）：

1. Go对各种变量、方法、函数等命名时使用的字符序列成为标识符
2. 凡是自己可以起名字的地方都叫标识符

标识符的命名规则：

1. 由26个英文字母大小写，0-9，_组成
2. 数字不可以开头
3. Gloang中严格区分大小写
4. 标识符不能含空格
5. 下划线_本身在Go中是一个特殊的标识符，称为空标识符。可以代表任何其他的标识符，但是它对应的值会被忽略（比如：忽略某个返回值）。所以**仅能作为占位符**使用，**不能作为标识符**使用
6. 不能以**系统保留关键字**作为标识符，比如break if等

###### 标识符命名注意事项：

1. 包名：保持package的名字和目录保持一致，尽量采取有意义的包名，简短有意义，不要和标准库冲突
2. 变量名、函数名、常量名尽量采用驼峰法
3. 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用（注：可以简单的理解成，**首字母大写是公有的，首字母小写是私有的**），在golang没有public，private等关键字

###### 系统保留关键字：

为了简化代码编译过程中对代码的解析，其定义的保留关键字只有25个，详见如下

| break    | default      | func   | interface | select |
| -------- | ------------ | ------ | --------- | ------ |
| case     | defer        | go     | map       | struct |
| chan     | else         | goto   | package   | switch |
| const    | fallthrought | if     | range     | type   |
| continue | for          | import | return    | var    |

除了保留关键字还有36个预定义标识符，其中包括基础数据类型和系统内嵌函数

###### 运算符介绍：

可以表示数据的运算、赋值和比较等

- 算数运算符
- 赋值运算符
- 比较运算符/关系运算符
- 逻辑运算符
- 位运算符
- 其他运算符& *

go里面没有三元运算符

###### 算数运算符：

如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分

如果希望保留小数部分，则需要有浮点数参与运算

a % b = a - a / b * b

算数运算符细节：

- 对于除号/，它的整数除和小数除是有区别的：整数之间做除法时，只保留整数部分而舍弃小数部分。例如 x:=19/5，结果是3

- 当对一个数取模时，可以等价a % b = a - a / b * b，这样我们可以看到取模的一个本质运算

- Go的**自增自减只能当做一个独立语言使用**时，不能这样使用b:=a++或者b:=a--

  ```go
  // 只能这样用
  i++
  a = i
  ```

- Go的++和--只能写在变量的后面，不能写在变量的前面，即只有a++  a-- 没有 ++a  --a

- Go的设计者去掉 c / java 中的自增自减的容易混淆的写法，让Go更加简介统一（强制性的）

###### 关系运算符：









