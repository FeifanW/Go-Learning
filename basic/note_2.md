##### 函数：

为完成某一功能的程序指令（语句）的集合，称为函数

在Go中函数分为**自定义函数**、**系统函数**

基本语法：

```go
func 函数名 (形参列表) (返回值类型列表) {
    执行语句...
    return 返回值列表
}
```

1. 形参列表：表示函数的输入
2. 函数中的语句：表示为了实现某一功能代码块
3. 函数可以有返回值，也可以没有

##### 包：

utils.go //专门用于定义函数，让其它文件来调用

db.go // 专门定义对数据库的操作的函数

包的本质就是创建不同的文件夹，来存放程序文件

一个包往往对应一个文件夹

包的三大作用：

1. 区分相同名字的函数、变量等标识符
2. 当程序文件很多时，可以很好的管理项目
3. 控制函数、变量等访问范围，即作用域

包的相关说明：

- 打包基本语法

  package util

- 引入包的基本语法

  import "包的路径"  默认路径就是src/所以直接写路径就行

  GOPATH/src/完整包名

- 跨包使用必须大写，类似其他语言的public

- 调用的时候是包名.函数

包的使用细节：

1. 在给文件打包时吗，该包对应一个文件夹，比如utils文件夹对应的包名就是utils，文件的包名通常和文件所在的文件夹名一致，一般是小写字母

2. 当一个文件要使用其他包函数或变量时，需要引入对应的包

3. 引入方式1：import "包名"

4. 引入方式2：

   import (

   ​	"包名"

   ​	"包名"

   )

5. package 指令在文件第一行，然后是import指令

6. 在import包时，路径从$GOPATH的src下开始，不用带src，编译器会自动从src下开始引入

7. 为了让其他包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其他语言的public，这样才能跨包访问，比如utils.go

8. 在访问其他包函数时，其语法是 包名.函数名，比如这里的main.go文件中

9. 如果包名比较长，Go支持给包取名，注意细节：取别名后，原来的包名就不能使用了

   ```go
   import (
   	"fmt"
       util "go_code/utils"   // 起别名
   )
   ```

10. 在同一包下，不能有相同的函数名（也不能有相同的全局变量名），否则报重复定义

11. 如果你要编译成一个可执行程序文件，就需要将这个包声明为main，即package main 这个就是一个语法规范，如果你是写一个库，包名可以自定义

    说明：

    - 演示一个案例，项目的目录如右图
    - 编译的指令，在项目目录下，编译路径不需要带src，编译器会自动带
    - 编译时需要编译main包所在的文件夹
    - 项目的目录结构最好按照规范来组织
    - 编译后生成一个有默认名的可执行文件，在$GOPATH目录下，可以指定名字和目录，比如：放在bin目录下 D:\goproject>go build -o bin/my.exe go_code/project/main

##### 函数调用机制底层剖析：

栈区：（基本数据类型一般分配到栈区，编译器存在一个逃逸分析）

堆区：（引用数据类型一般说分配到堆区，编译器存在一个逃逸分析）

代码区：代码存放到这

说明：

- 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间和其他栈的空间区分开
- 在每个函数对应的栈中，数据空间是空间的，不会混淆
- 当一个函数调用完毕（执行完毕）后，程序会销毁这个函数对应的栈空间

return语句：

```go
func 函数名（形参列表）（返回值类型列表）{
    语句...
    return 返回值列表
}
```

- 如果返回多个值，在接收时，希望忽略某个返回值，则使用_符号表示占位忽略
- 如果返回值只有一个（返回值类型列表）可以不写()

```go
res1,res2 := getSumAndSub(1,2)
// 如果只想获取一个值，可以这么做
_, res3 := getSumAndSub(1,2)
```

递归调用：

函数在函数体内调用了本身

- 递归必须向退出递归条件逼近，否则就是无限循环调用
- 执行一个函数时，就创建一个新的受保护的独立空间（新函数栈）
- 函数的局部变量是独立的，不会相互影响
- 当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或返回时，该函数本身也会被系统摧毁

函数注意事项和细节：

1. 函数的形参列表可以是多个，返回值列表也可以是多个

2. 形参列表和返回值列表的数据类型可以是值类型和引用类型

3. 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其他包文件使用，类似public，首字母小写，只能被本包文件使用，其他包文件不能使用，类似private

4. 函数中的变量是局部的，函数外不生效

5. 基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值

6. 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量

7. Go函数不支持重载

8. 在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量，通过该变量可以对函数调用

9. 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用！ 

10. 为了简化数据类型定义，Go支持自定义数据类型

   基本语法：type 自定义数据类型名 数据类型  // 理解：相当于一个别名

   案例：type myInt int // 这是myInt就等价int来使用了

   案例：type mySum func(int，int) int  // 这时mySum就等价一个函数类型func (int，int) int

11. 支持对函数返回值命名

    ```go
    func cal(n1 int, n2 int) (sum int, sub int) {
        sum = n1 + n2
        sub = n1 - n2
        return
    }
    ```

12. 使用 _ 标识符，忽略返回值

13. Go支持可变参数

    ```go
    // 支持0到多个参数
    func sum(args... int) sum int {
    }
    // 支持1到多个参数
    func sum(n1 int, args... int) sum int {
    }
    ```

    说明：

    1. args是slice切片，通过args[index]可以访问到各个值
    2. 如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后

##### init函数：

每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，也就是说init会在main函数前被调用

细节：

1. 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是

   变量定义 -> init函数 -> main函数

2. init函数最主要的作用，就是完成一些初始化的工作，比如下面的案例

   比如初始化变量

##### 匿名函数：

Go支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数可以实现多次调用

- 使用方式1：

  在定义匿名函数时就直接调用

  ```go
  res1 := func (n1 int, n2 int) int {
      return n1 + n2
  }(10, 20)
  ```

- 使用方式2：

  将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数

  ```go
  a := func (n1 int, n2 int) int {
      return n1 - n2
  }
  res2 := a(10, 30)
  ```

- 全局匿名函数：

  ```go
  var (
  	// fun1就是一个全局匿名函数
      Fun1 = func (n1 int, n2 int) int {
          return n1 * n2
      }
  )
  ```

##### 闭包：

闭包就是一个函数和其相关的引用环境组合的一个整体（实体）

和js里的闭包有点像哈哈哈

```go
// 累加器
func AddUpper() func (int) int{
    var n int = 10
    return func (x int) int {
        n = n + x
        return n
    }
}

func main() {
    // 使用前面的代码
    f := AddUpper()
    fmt.Println(f(1))  // 11
    fmt.Println(f(1))  // 13
    fmt.Println(f(3))  // 16
}
```

说明和总结：

1. AddUpper是一个函数，返回的数据类型是fun (int) int

2. 闭包的说明

   返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包

3. 可以理解为：闭包是类，函数是操作，n是字段。函数和它使用到n构成闭包

4. 当我们反复的调用f函数时，因为n是初始化一次，因此每调用一次就进行累计

5. 我们要搞清楚闭包的关键，就是分析出返回的函数它使用（引用）到那些变量，因为函数和它引用到的变量共同构成闭包

闭包案例：

1. 编写一个函数makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包

2. 调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如.jpg），则返回文件名.jpg，如果已经有.jpg后缀，则返回原文件名

3. 要求使用闭包的方式完成

4. strings.HasSuffix，该函数可以判断某个字符串是否有指定的后缀

   ```go
   func makeSuffix(suffix string) func (string) string {
       return func (name string) string {
           // 如果 name 没有指定后缀，则加上，否则就返回原来的名字
           if !strings.HasSuffix(name, suffix) {
               return name + suffix
           }
           return name
       }
   }
   ```

说明：

1. 返回的函数和makeSuffix (suffix string) 和 suffix变量 和返回的函数组合成一个闭包，因为返回的函数引用到suffix这个变量
2. 体会一下闭包的好处，如果用传统的方法，也可以实现这个功能，但是传统方法需要每次都传入后缀名，比如.jpg，而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用

##### 函数defer

在函数中经常需要创建资源（比如：数据库连接、文件句柄、锁等），为了在函数执行完毕后，及时释放资源，Go的设计者提供defer（延迟机制）

```go
func sum(n1 int, n2 int) int {
    // 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
    // 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
    defer fmt.Println("ok1 n1=",n1)  // 3
    defer fmt.Println("ok2 n2=",n2)  // 2
    res := n1 + n2  
    fmt.Println("ok3 res=",res)   // 1
    return res
}
func main(){
    res := sum(10,20)
    fmt.Println("res=",res)  // 4
}
```

细节：

1. 当Go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数的下一个语句
2. 当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行（注：遵守栈 先入后出的机制），所以同学们看到前面的案例输出的顺序
3. 在defer将语句放入到栈时，也会将相关的值拷贝同时入栈

defer最主要的价值在，当函数执行完毕后，可以及时释放函数创建的资源

```go
func test() {
    // 关闭文件资源
    file.openfile(文件名)
    defer file.close()
    // 其他代码
}
```

```go
func test() {
    // 释放数据库资源
    connect = openDatabse()
    defer connect.close()
    // 其他代码
}
```

说明：

1. 在Go编程中的通常做法是，创建资源后，比如（打开了文件，获取了数据库的链接，或者是锁资源），可以执行 defer.file.Close() defer connect.Close()
2. 在defer后，可以继续使用创建资源
3. 当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源
4. 这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心

##### 函数参数传递方式：

两种传递方式：

1. 值传递
2. 引用传递

其实不管值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是地址的拷贝，一般来说，地址拷贝效率高，因为数据量小，而值拷贝决定拷贝的数据大小，数据越大，效率越低

值类型和引用类型：

1. 值类型：基本数据类型 int 系列、float系列、bool、string、数组和结构体struct
2. 引用类型：指针、slice切片、map、管道chan、interface 等都是引用传递

##### 变量作用域：

1. 函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部
2. 函数外部声明/定义的变量叫全局变量，作用域在整个包都有效，如果其首字母为大写，则作用域在整个程序有效
3. 如果变量是在一个代码块，比如for/if中，那么这个变量的作用域就在该代码块

```go
Name := "tom"  // var Name string  Name = "tom"会报错因为赋值语句只能在函数体内，函数外可以初始化不能赋值
```

##### Go字符串函数：

1. 统计字符串的长度，按字节len(str)
2. 字符串遍历，同时处理有中文的问题 r:=[]rune(str)
3. 字符串转整数：n, err := strconv.Atoi("12")
4. 整数转字符串 str = strconv.Itoa(123456)
5. 字符串 转[]byte: var bytes = []byte("hello go")
6. []byte 转 字符串 :str = string([]byte{97,98,99})
7. 10进制转2,8,16进制 :str=strconv.FormatInt(123,2)  // 1 -> 8, 16
8. 查找子串是否在指定进制的字符串中 :strings.Contains("seafood","foo")  // true
9. 统计一个字符串有几个指定的子串  :strings.Count("ceheese","e")  // 4
10. 不区分大小写的字符串比较（==是区分字母大小写的）:fmt.Println(strings.EqualFold("abc","Abc"))  // true
11. 返回子串在字符串中第一次出现的index值，如果没有返回-1 :strings.Index("NLT_abc","abc")   // 4
12. 返回子串在字符串最后一次出现的index，如果没有返回-1：strings.LastIndex("go golang","go")
13. 将指定的子串替换成另一个子串：strings.Replace("go go hello","go","go语言",n) n 可以指定你希望替换几个，如果n=-1表示全部替换
14. 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,world",",")
15. 将字符串的字母进行大小写的转换：strings.ToLower("Go")  // go strings.ToUpper("Go")  // Go
16. 将字符串左右两边的空格去掉：strings.TrimSpace("  javascript     ")
17. 将字符串左右两边指定的字符去掉：strings.Trim("!hello! "," !")   将左右两边的！和" "去掉
18. 将字符串左边指定的字符去掉：strings.TrimLeft("! hello!","!")   将左边的！和" "去掉
19. 将字符串右边指定的字符去掉：strings.TrimRight("! hello!","!")  将右边的！和" "去掉
20. 判断字符串是否以指定的字符串开头 strings.HasPrefix("ftp://192.168.10.1","ftp")  
21. 判断字符串是否以指定的字符串结束 strings.HasSuffix("NLT_abc.jpg","abc")

```go
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
```

##### Go的时间和日期函数详解：

需要导入time包

1. 获取当前时间

   ```go
   now := time.Now()
   ```

2. 通过now可以获取到年月日，时分秒

   ```go
   now.Year()
   now.Month()  // 默认是英文 int(now.Month())可以转成数字
   now.Day()
   now.Hour()
   now.Minute()
   now.Second()
   ```

3. 格式化日期时间

   - 格式化的第一种方式

     ```go
     fmt.Printf("当前的年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
     
     dateStr := fmt.Sprintf("当前的年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
     
     fmt.Println("dateStr=%v",dateStr)
     ```

   - 格式化的第二种方式

     ```go
     // time提供的Format函数
     fmt.Printf(now.Format("2006-01-02 12:04:05"))
     fmt.Println()
     fmt.Printf(now.Format("2006-01-02"))
     fmt.Println()
     fmt.Printf(now.Format("12:04:05"))
     fmt.Println()
     ```

     说明：

     "2006/01/02 15:04:05" 这个字符串的各个数字是固定的，必须这样写

     "2006/01/02 15:04:05" 这个字符串各个数字可以自由的组合，这样可以按照程序需求来返回时间和日期

4. 时间的常量

   ```go
   const (
   	Nanosecond Duration = 1  //纳秒
       Microsecond = 1000 * Nanosecond  // 微秒
       Millisecond = 1000 * Microsecond  // 毫秒
       Second = 1000 * Millisecond  // 秒
       Minute = 60 * Second  // 分钟
       Hour = 60 * Minute  // 小时
   )
   常量的作用：在程序中可用于获取指定时间单位的时间，比如想得到100毫秒
   100 * time.Millisecond
   
   // 需求：每隔0.1秒打印一个数字，打印到100时就退出
   i := 0
   for {
       i++
       fmt.Println(i)
       // 休眠
       // time.Sleep(time.Millisecond * 100)
       if i == 100 {
           break
       }
   }
   ```

5. 获取当前unix时间戳 和 unixnano 时间戳 （作用是可以获取随机数字）

   unix时间戳：返回从1970年UTC到时间t所经过的时间（单位秒）

   unixnano时间戳：返回从1970年UTC到时间t所经过的时间（单位纳秒）

   ```go
   fmt.Printf("unix时间戳=%v unixnano时间戳=%v",now.Unix(), now.UnixNano())
   ```

##### 内置函数：

Go的设计者为了编程方便，提供了一些函数，这些函数可以直接使用，称之为Go的内置函数

1. len：用来长度，比如string、array、slice、map、channel

2. new：用来分配内存，主要来分配值类型，比如int、float32、struct...返回的是指针

   ```go
   num2 := new(int)  // *int
   /*
   num2的类型%T => *int
   num2的值 = 地址 exc04204c098(这个地址是系统分配)
   num2的地址%v = 地址 exc04206a020(这个地址是系统分配)
   num2指向的值 = 100
   */
   *num2 = 100
   fmt.Printf("num2的类型%T,num2的值=%v,num2的地址%v\n num2这个指针，指向的值=%v",num2,num2,&num2,*num2)
   ```

3. make：用来分配内存，主要用来分配引用类型，比如channel、map、slice

##### Go错误处理机制：

1. 默认情况下，当发生错误后（panic），程序就会退出（崩溃）
2. 如果我们希望，当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行，还可以在捕获到错误后，给管理员一个提示（邮件，短信...）

基本说明：

1. Go语言追求简洁优雅，所以Go不支持传统的try...catch...finally这种处理
2. Go中引入的处理方式为：defer、panic、recover
3. 在这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过**recover捕获这个异常**，然后正常处理

```go
fun test(){
    // 使用defer + recover 来捕获和处理异常
    defer func(){
        err := recover()  //recover()内置函数，可以捕获到异常
        if err != nil {  // 捕获到错误
            fmt.Println("err=",err)
            // 这里就可以将错误信息发送给管理员...
        }
        
        /*
        另外一种写法
        if err := recover();  err != nil {
        
        }
        */
    }()
}
```

错误处理的好处：

进行错误处理后，程序不会轻易的挂掉，如果加入预警代码，就可以让程序更加的健壮

###### 自定义错误：

Go程序中，也支持自定义错误，使用errors.New 和 panic 内置函数

1. errors.New("错误说明")  会返回一个error类型的值，表示一个错误
2. panic内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数，可以接收error类型的变量，**输出错误信息，并退出程序**

```go
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
```

##### 数组和切片：

数组可以存放多个同一类型数据，数组也是一种数据类型，**在Go中数组是值类型**

```go
// 定义一个数组
var hens [6]float64;
...
var avg = totalWeight / float64(len(hens))  // 当分母写变量时，类型需要和分子一致
var avg = totalWeight / 6  // 当分母是数字时则不需要转 
```

定义：

var 数组名 [数组大小]数据类型

var a [5]int   //  int占8个字节

赋初值 a[0] = 1 a[1] = 30...

总结：

1. 数组的地址可以通过数组名来获取 &intArr
2. 数组的第一个元素的地址，就是数组的首地址
3. 数组的各个元素的地址间隔是依据数组的类型决定int64 -> 8 int32 ->4...

###### 访问数组元素：

数组名 [下标]

```go
var score [5]float64
//用户输入的方式给数组赋值
fmt.Scanln(&score[i])
```

四种初始化数组的方式：

```go
var numArr01 [3]int = [3]int{1,2,3}
var numArr02 = [3]int{5,6,7}
var numArr03 = [...]int{8,9,10}
var numArr04 = [...]int{1:800, 0:900, 2:791}  // 下标
numArr05 := [...]string{1:"tom", 0:"jack", 2:"mary"}  // 类型推导也可以
```

###### 数组的遍历：

1. 常规遍历

2. for-range结构遍历

   Go语言独有的结构，可以用来遍历访问数组的元素

   ```go
   for index,value := range array01 {
       ...
   }
   ```

   说明：

   1. 第一个返回值index是数组的下标
   2. 第二个value是在该下标位置的值
   3. 他们都是仅在for循环内部可见的局部变量
   4. 遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标位下划线_
   5. index和value的名称是不固定的，即程序员可以自行指定，一般命名为index和value

   ```go
   for i,v := range heros {
       
   }
   for _,v := range heros {
       
   }
   ```

数组细节：

1. 数组是多个相同类型数据的组合，一个数组一旦声明/定义，**长度是固定的，不能动态变化**

2. var arr []int 这是arr就是一个slice切片

3. 数组中元素可以是任何数据类型，包括值类型和引用类型，但是不能混用

4. 数组创建后，如果没有赋值，有默认值

   数值类型数组 默认为0

   字符串数组  默认为""

   bool数组  默认值为false

   ```go
   var arr01 [3]float32
   var arr02 [3]string
   var arr03 [3]bool
   fmt.Printf("%v,%v,%v",arr01,arr02,arr03)
   ```

5. 数组下班必须在指定范围内使用，否则报panic：数组越界，比如var arr[5]int 则有效下标为0-4

6. Go的数组属值类型，在默认情况下是值传递，因为会进行值拷贝。数组间不会相互影响

   ```go
   func test01(arr [3]int) {
       arr[0] = 88
   }
   
   arr := [3]int{11,22,33}
   test01(arr)
   fmt.Println(arr)   // arr[0]并没有变成88
   ```

7. 如果在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）

   ```go
   func test01(arr *[3]int) {
       (*arr)[0] = 88
   }
   
   arr := [3]int{11,22,33}
   test01(&arr)
   fmt.Println(arr)   // arr[0]会被修改
   ```

8. 长度是数组类型的一部分，在传递函数参数时，**需要考虑数组的长度**

9. 不能把数组类型传递给切片

###### 切片：

1. 切片的英文是slice

2. 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制

3. 切片的使用和数组类似，遍历切片、访问切片的冤死和求切片长度len都一样

4. **切片的长度是可以变化的**，因此切片是一个**可以动态变化数组**

5. 切片的定义的基本语法：

   var 变量名 []类型

   ```go
   var intArr [5]int = [...]int{1,22,33,66,99}
   /*
   声明/定义一个切片
   1.slice就是切片名
   2.intArr[1:3]表示slice引用到intArr这个数组
   3.引用intArr数组的起始下标为1，最后的下标为3（但是不包含3）
   */
   slice := intArr[1:3]
   fmt.Println("intArr=",intArr)
   fmt.Println("slice 的元素是=",slice)
   fmt.Println("slice 的元素个数",len(slice))
   fmt.Println("slice 的容量 =",cap(slice))
   ```

   slice在内存中可以理解为由三部分组成：

   - 第一个部分是引用的第一个数组元素的地址
   - 第二个是slice本身的长度
   - 第三个是slice容量的大小

   总结：

   从上面可以看到

   1. slice的确是一个引用类型

   2. slice从底层来说，其实就是一个数据结构（struct结构体）

      ```go
      type slice struct{
          ptr *[2]int
          len int
          cap int
      }
      ```

###### 切片的使用：

- 方式1：定义一个切片，然后让切片去引用一个已经创建好的数组，比如前面的案例

- 方式2：通过make来创建切片

  基本语法：var 切片名 []type = make([],len,[cap])

  参数说明：type就是数据类型len：大小 cap：指定切片容量，可选

  切片默认值都为0

  ```go
  var slice []float64 = make([]float64, 5, 10)
  slice[1] = 10
  slice[3] = 20
  fmt.Println(slice)
  fmt.Println("slice的size=", len(slice))
  fmt.Println("slice的cap=", cap(slice))
  ```
  
  总结：
  
  1. 通过make方式创建切片可以指定切片的大小和容量
  2. 如果没有给切片的各个元素赋值，那么就会使用默认值
     - int, float => 0
     - string => ""
     - bool => false
  3. 通过make方式创建的切片对应的数组是由make底层维护，对外不可即，只能通过slice去访问各个元素
  
- 方式3：定义一个切片，直接就指定具体数组，使用原理类似make的方式

  ```go
  var strSlice []string = []string{"tom", "jack", "mary"}
  fmt.Println("strSlice=", strSlice)
  fmt.Println("strSlice size=", len(strSlice))
  fmt.Println("strSlice=", cap(strSlice))
  ```

###### 切片的遍历：

- for循环常规方式遍历

- for-range 结构遍历切片

  ```go
  //使用常规的for循环遍历切片
  var arr [5]int = [...]int{10, 20, 30, 40, 50}
  slice := arr[1:4]
  for i := 0; i < len(slice); i++ {
      fmt.Printf("slice[%v]=%v", i, slice[i])
  }
  fmt.Println()
  //使用for-range方式遍历切片
  for i, v := range slice {
      fmt.Printf("i=%v v=%v \n", i, v)
  }
  ```

说明：

1. 切片初始化时 var slice = arr[startIndex:endIndex] 左闭右开

2. 切片初始化时，仍然不能越界，范围在[0-len(arr)]之间，但是可以动态增长

   - var slice = arr[0:end]  可以简写 var slice = arr[:end]
   - var slice = arr[start len(arr)]可以简写 var slice = arr[start:]
   - var slice = arr[0:len(arr)] 可以简写: var slice = arr[:]

3. cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素

4. 切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组或make一个空间供切片来使用

5. 切片可以继续切片

   ```go
   slice2 := slice[1:2]   // 指向的是相同的空间
   ```

6. 用append内置函数，可以对切片进行动态增加

   ```go
   var arr [5]int = [...]int{10, 20, 30, 40, 50}
   slice := arr[1:4]
   //用append内置函数，可以对切片进行动态增加
   var slice3 []int = []int{100, 200, 300}
   // 通过append直接给slice3追加具体的元素
   slice3 = append(slice3, 400, 500, 600)
   fmt.Println("slice3", slice3)
   
   // 通过append将切片slice3追加给slice3
   slice3 = append(slice3, slice...)
   fmt.Println("slice3", slice3)
   ```

   切片append操作的底层原理分析：

   1. 切片append操作的本质就是对数组扩容
   2. go底层会创建一个新的数组newArr(安装扩容后大小)
   3. 将slice原来包含的元素拷贝到新的数组newArr
   4. slice重新引用到newArr
   5. 注意newArr是在底层来维护的，程序员不可见

7. 切片的拷贝操作

   切片使用copy内置函数完成拷贝，举例说明

   ```go
   var arr [5]int = [...]int{10, 20, 30, 40, 50}
   slice := arr[1:4]
   ```

   说明：

   1. copy(para1,para2)参数的数据类型是切片
   2. 按照上面的代码来看，slice4和slice5的数据空间是独立，相互不影响，也就是说slice[0]=999 slice5[0] 仍然是1
   3. 拷贝不会扩容

8. 切片是引用类型，所以在传递时，遵循引用传递机制

###### string和slice:

1. string底层是一个byte数组，因此string也可以进行切片处理

   ```go
   str := "hello world"
   slice := str[6:]
   fmt.Println("slice=",slice)
   ```

2. string和切片在内存的形式，以"abcd"画出内存示意图

3. string是不可变的，也就是说不能通过str[0] = 'z' 方式来修改字符串

4. 如果需要修改字符串，可以先将string -> []byte / 或者 []rune -> 修改 -> 重写转成string

   ```go
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
   ```

##### 排序和查找：

1. 内部排序

   将需要处理的所有数据都加载到内部存储器中进行排序

   包括（交换式排序法、选择式排序法和插入式排序法）

2. 外部排序

   数据量过大，无法全部加载到内存中，需要借助外部存储进行排序，包括（合并排序法和直接合并排序法）

- 交换式排序法

  - 冒泡排序

    ```go
    //冒泡排序
    func BubbleSort(arr *[5]int) {
    	fmt.Println("排序当前arr=", (*arr))
    	temp := 0
    	for i := 0; i < len(*arr)-1; i++ {
    		for j := 0; j < len(*arr)-1-i; j++ {
    			if (*arr)[j] > (*arr)[j+1] {
    				// 交换
    				temp = (*arr)[j]
    				(*arr)[j] = (*arr)[j+1]
    				(*arr)[j+1] = temp
    			}
    		}
    	}
    	fmt.Println("排序后的arr=", (*arr))
    }
    
    func main() {
    	arr := [5]int{24, 68, 80, 57, 13}
    	BubbleSort(&arr)
    }
    ```
  
  - 快速排序
  
  ###### 二分查找：
  
  ```go
  // 二分查找
  func BinaryFind(arr *[5]int, leftIndex int, rightIndex int, findVal int) {
     // 判断leftIndex是否大于rightIndex
     if leftIndex > rightIndex {
        fmt.Println("找不到")
        return
     }
     // 先找到 中间的下标
     middle := (leftIndex + rightIndex) / 2
     if (*arr)[middle] > findVal {
        BinaryFind(arr, leftIndex, middle-1, findVal)
     } else if (*arr)[middle] < findVal {
        BinaryFind(arr, middle+1, rightIndex, findVal)
     } else {
        fmt.Printf("找到了，下标为%v\n", middle)
     }
  }
  arr := [5]int{24, 68, 80, 57, 13}
  BinaryFind(&arr, 0, len(arr)-1, 1000)
  ```

##### 二维数组：

- 使用方式1：var 数组名 \[大小][大小]类型

- 使用方式2：var 数组名 \[大小][大小]类型 = \[大小][大小]类型{{初值..},{初值..}}

  赋值(有默认值，比如int 类型的就是0)

说明：

二维数组在声明/定义时也有对应的四种写法[和一维数组类似]

```go
var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值...},{初值...}}
var 数组名 [大小][大小]类型 = [...][大小]类型{{初值..},{初值..}}
var 数组名 = [大小][大小]类型{{初值..},{初值..}}
var 数组名 = [...][大小]类型{{初值..},{初值..}}
```

###### 二维数组的遍历：

```go
// 演示二维数组的遍历
var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

// for循环来遍历
for i := 0; i < len(arr3); i++ {
    for j := 0; j < len(arr3[i]); j++ {
        fmt.Printf("%v\t", arr3[i][j])
    }
    fmt.Println()
}

// for-range来遍历二维数组
for i, v := range arr3 {
    for j, v2 := range v {
        fmt.Printf("arr3[%v][%v]=%v \t", i, j, v2)
    }
    fmt.Println()
}
```

代码太长换行的时候末尾带一个逗号

##### map:

map是key-value数据结构，又称为字段或关联数组，类似其它编程语言的集合在编程中经常使用到

基本语法：

var map 变量名 map[keytype]valuetype

- key可以是什么类型

  golang中的map的key可以是多种类型，比如bool、数字、string、指针、channel还可以是只包含前面几个类型的接口、结构体、数组。通常为**int、string**

  注意：slice、map还有function不可以，因为这几个没法用 == 来判断

- valuetype可以是什么类型

  valuetype的类型和key基本一样

  通常为：数字（整数、浮点数）string、map、struct

map声明举例：

```go
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
```

注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用

```go
// map的声明和注意事项
var a map[string]string
// 组使用map前，需要先make，make的作用就是给map分配数据空间
a = make(map[string]string, 10)
a["1"] = "宋江"
a["2"] = "吴用"
a["3"] = "武松"
a["4"] = "公孙胜"
fmt.Println(a)
```

说明：

1. map在使用前一定要make
2. map的key是不能重复，如果冲服，则以最后这个key-value为准
3. map的value是可以相同的
4. map的key-value 是无序的
5. make内置函数数目

###### map的使用方式：

- 方式1：

  ```go
  // map的声明和注意事项
  var a map[string]string
  // 组使用map前，需要先make，make的作用就是给map分配数据空间
  a = make(map[string]string, 10)
  a["1"] = "宋江"
  a["2"] = "吴用"
  a["3"] = "武松"
  a["4"] = "公孙胜"
  fmt.Println(a)
  ```

- 方式2：

  ```go
  // 第二种方式
  cities := make(map[string]string)
  cities["1"] = "北京"
  cities["2"] = "天津"
  cities["3"] = "上海"
  fmt.Println(cities)
  ```

- 方式3：

  ```go
  // 第三种方式
  heroes := map[string]string{
  "hero1": "宋江",
  "hero2": "卢俊义",
  }
  fmt.Println(heroes)
  ```

案例：

```go
/*
练习：演示一个key-value 的value是map的案例
比如：我们要存放3个学生信息，每个学生有name和sex信息
思路：map[string]map[string]string
*/
studentMap := make(map[string]map[string]string)

studentMap["stu01"] = make(map[string]string, 3)
studentMap["stu01"]["name"] = "tom"
studentMap["stu01"]["sex"] = "男"
studentMap["stu01"]["address"] = "北京长安街"

studentMap["stu02"] = make(map[string]string, 3) // 不能少
studentMap["stu02"]["name"] = "mary"
studentMap["stu02"]["sex"] = "女"
studentMap["stu02"]["address"] = "上海黄浦江"

fmt.Println(studentMap)
fmt.Println(studentMap["stu02"])
fmt.Println(studentMap["stu02"]["address"])
```

###### map的增删改查crud操作：

map增加和更新：

map["key"] = value  // 如果key还没有，就是增加，如果key存在就是修改

```go
cities := make(map[string]string)
cities["1"] = "北京"
cities["2"] = "天津"
cities["3"] = "上海"
fmt.Println(cities)
cities["3"] = "上海~"
fmt.Println(cities)
```

map删除：

delete(map,"key")，delete是一个内置函数，如果key存在，就删除该key-value如果key不存在，不操作，但是也不会报错

```
delete(cities,"1")
// 当delete指定的key不存在时，删除不会操作，也不会报错
```

细节说明：

- 如果要删除map所以的key，没有一个专门的方法一次删除，可以遍历一下key逐个删除

- 或者map = make(...)，make一个新的，让原来的称为垃圾，被gc回收

  ```go
  cities = make(map[string]string)
  fmt.Println(cities)
  ```

map查找：

```go
val,findRes = heroes["no1"]
```

说明：如果heroes这个map中存在"no1"，那么findRes就会返回true 否则返回false

演示：

```go
val,ok := cities["no2"]
if ok {
    fmt.Printf("有no1 key值为%v\n", val)
} else {
    fmt.Printf("没有no1 key\n")
}
```

map遍历：

使用for-range的结构遍历

```go
// for-range遍历
for k, v := range cities {
	fmt.Printf("k=%v v=%v", k, v)
}
```

map的长度：

fmt.Println(len(stus))

###### map切片：

map的个数可以动态变化

```go
//这里我们需要使用到切片的append函数，可以动态的增加monster
//1.先定义个monster信息
newMonster := map[string]string{
    "name": "新的妖怪",
    "age":  "200",
}
monsters = append(monsters, newMonster)
fmt.Println(monsters)
```

###### map排序：

- Go中没有一个专门的方法针对map的key进行排序
- Go中的map默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历得到的输出可能不一样
- Go中map的排序，是先将key进行排序，然后根据key值遍历输出即可

```go
//map的排序
map1 := make(map[int]int, 10)
map1[10] = 100
map1[1] = 13
map1[4] = 56
map1[8] = 90
fmt.Println(map1)
// 如果按照map的key的顺序进行排序输出
//1.先将map的key放入到切片中
//2.对切片排序
//3.遍历切片，然后按照key来输出map的值
var keys []int
for k, _ := range map1 {
    keys = append(keys, k)
}
// 排序
sort.Ints(keys)
fmt.Println(keys)

for _, k := range keys {
    fmt.Printf("map1[%v]=%v \n", k, map1[k])
}
```

###### map使用细节：

- map是引用类型，遵循引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
- map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能**动态的增长键值对**（key-value）
- map的value也经常使用struct类型， 更适合管理复杂的数据（比前面的value是一个map更好），比如value为Student结构体

```go
package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
}

//定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// map是引用类型，遵守引用类型传递的机制，在一个函数接收map
	// 修改后，会直接修改原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	fmt.Println(map1)

	//map的value 也经常使用struct类型
	//更适合管理复杂的数据（比前面的value是一个map更好）
	//比如value为Student结构体
	//1.map的key为学生的学号，是唯一的
	//2.map的value为结构体，包含学生的姓名，年龄，地址
	students := make(map[string]Stu, 10)
	// 创建2个学生
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 20, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)
	// 遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
		fmt.Println()
	}
}

```

案例：

1. 使用map[string]map[string]string 的map类型
2. key：表示用户名，是唯一的，不可以重复
3. 如果某个用户名存在，就将其密码修改为"888888"，如果不存在就增加这个用户信息，(包括昵称nickname)
4. 编写一个函数 modifyUser(users  map[string]map[string]string, name string) 完成上述功能

##### 面向对象编程：

说明：

- Go支持面向对象编程（OOP），但是和传统的面向对象编程有区别，并不是**纯粹的面向对象的语言**，所以说**Go支持面向对象编程特性**是比较准确的
- Go没有类class，Go语言的结构体(struct)和其它编程语言的类(class)有同等的地位，你可以理解Go是基于struct来实现OOP特性的
- Go面向对象编程非常简洁，去掉了传统OOP语言的继承、方法重载、构造函数和析构函数、隐藏的this指针等等
- Go仍然有面向对象编程的继承、封装和多态的特性，只是实现的方式和其他OOP语言不一样，**比如继承**：Go没有extends关键字，继承是通过匿名字段来实现
- Go的面向对象(OOP)很优雅，OOP本身就是语言类型系统的一部分，通过接口（interface）关联，耦合性低，也非常灵活

```go
package main

import "fmt"

//定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	// 创建一个Cat变量
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"
	fmt.Println("cat1=", cat1)
	fmt.Println("猫猫的信息如下", "")
	fmt.Println("Age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)

}
```

结构体和结构体变量（实例）的区别和联系

1. 结构体是自定义的数据类型，代表一类事物
2. 结构体变量（实例）是具体的，实际的，代表一个具体变量

结构体是值类型，不是引用类型

声明结构体：

```go
type 标识符 struct{
	field1 type
	field2 type
}
```

1. 从概念或叫法上看：结构体字段 = 属性 = field (即授课中：统一叫字段)
2. 字段是结构体的一个组成部分，一般是基本数据类型、数组，也可以是引用类型。比如我们前面定义的猫结构体的Name string 就是属性

```go
type Cat struct {
	Name  string
	Age   int
	Color string
	Scores [3]int   // 可以定义数组
	Hobby string
}
```

字段/属性

注意事项和细节说明：

1. 字段声明语法同变量，示例：字段名 字段类型

2. 字段的类型可以为：基本类型、数组或引用类型

3. 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前将的一样：

   布尔类型是false 数字是0 字符串是""

   数组类型的默认值和它的元素类型相关，比如score [3]int则为[0,0,0]

   指针 slice 和 map的零值都是nil，即还没有分配空间

4. **不同的结构体变量**的字段是独立的，互不影响，一个结构体变量字段的更改，不影响另外一个

```go
// 如果结构体的字段类型是：指针、slice和map的零值都是nil，即还没有分配空间
// 如果需要使用这样的字段，需要先make,才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string // 切片
}
type Monster struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string // 切片
}

func main() {
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	// 使用slice 再次说明，一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// 使用map,一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"
	fmt.Println(p1)

	//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改不影响另外一个，结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 // 结构体是值类型，默认为值拷贝
	//monster2 := &monster1 // 如果想改同一个，则传地址
	monster2.Name = "青牛精"

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
}
```

##### 创建结构体实例的四种方法：

```go
// 如果结构体的字段类型是：指针、slice和map的零值都是nil，即还没有分配空间
// 如果需要使用这样的字段，需要先make,才能使用
type Person struct {
	Name string
	Age  int
	//Scores [5]float64
	//ptr    *int
	//slice  []int
	//map1   map[string]string // 切片
}
func main() {
	// 方式1
	// 方式2
	p2 := Person{"mary", 20}
	p2.Name = "tom"
	p2.Age = 18
	fmt.Println(p2)

	// 方式3-&
	// 案例：var person *Person = new (Person)

	var p3 *Person = new(Person)
	// 因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith" 也可以这样写 p3.Name = "smith"
	//原因：go的设计者，为了程序员使用方便，底层会对p3.Name = "smith"进行处理
	// 会给p3加上 取值运算（*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "john"
	(*p3).Age = 30
	fmt.Println(*p3)

	//方式4 -{}
	// 案例：var person *Person = &Person{}
	// 下面的语句，也可以直接给字符串赋值
	// var person *Person = &Person{"mary",60}
	var person *Person = &Person{}
	// 因为person是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以person.Name = "scott"
	// 原因和上面一样，底层会对person.Name = "scott" 进行处理，会加上（*person）
	(*person).Name = "scott"
	person.Name = "scott~~"
	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)

}
```

##### 结构体使用细节：

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1有四个int，在内存中是连续分布
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址是=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p  \n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的
	// 但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	// 打印地址
	fmt.Printf("r2.leftUp 本身地址是=%p r2.rightDown 本身地址=%p  \n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp 指向地址是=%p r2.rightDown 指向地址=%p  \n", r2.leftUp, r2.rightDown)

}
```

结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段（名字、个数和类型）

```go
type A struct {
	Num int
}
type B struct {
	Num int
}

func main() {
	var a A
	var b B
	a = A(b) // 可以转换，但是有要求，就是结构体的字段要完全一样（包括：名字、个数和类型）
	fmt.Println(a, b)
}
```

结构体进行type重新定义（相当于取别名），Go认为是新的数据类型，但是相互间可以强转

```go
type integar int
func main(){
    var i interger = 10
    var j int = 20
    j = int(i)   // j=i是不正确的
    fmt.Println(i,j)
}
```

struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，场景的使用场景就是序列化和反序列化

```go
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//1.创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	//2.将monster变量序列化为json格式字串
	// json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
```

##### 方法：

Go中的方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），因此**自定义类型，都可以有方法**，而不仅仅是struct

方法的声明和调用

```go
type A struct {
	Num int 
}
func (a A)test() {
	fmt.Println(a.Num)
}
```

对上面语法的说明

1. func (a A) test() {} 表示A结构体有一方法，方法名为 test
2. (a A)体现 test 方法是和A类型绑定的

```go
type Person struct {
	Name string
}

// 给Person类型绑定一方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
}
```

总结：

1. test方法和Person类型绑定
2. test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
3. func (p Person) test() {} ... p 表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相思
4. p这个名字，由程序员指定，不是固定，比如修改成person也是可以的

###### 方法的调用和传参机制：

说明：

1. 在通过一个变量去调用方法时，其调用机制和函数一样
2. 不一样的地方时，变量调用方法时，该变量本身也会作为一个参数传递到方法（如果变量是值类型，则进行只拷贝，如果变量是引用类型，则进行值拷贝）

###### 方法的声明（定义）：

注意事项和细节：

- 结构体类型是值类型，在方法调用中，遵循值类型的传递机制，是值拷贝传递方式

- 如果希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理

- Go中的**方法作用在指定的数据类型上**的（即：**和指定的数据类型绑定**），因此**自定义类型，都可以有方法**，而不仅仅是struct，比如int，float32等都可以有方法

  ```go
  type integer int
  func (i integer) print() {
      fmt.Println("i=",i)
  }
  //编写一个方法，可以改变i的值
  func (i *integer) change(){
      *i = *i + 1
  }
  func main() {
      var i integer = 10
      i.print()
      i.change()
      fmt.Println("i=",i)
  }
  ```

- 放大的访问范围控制的规则，和函数一样，方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其它包访问

- 如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出

  ```go
  type Student struct {
      Name string
      Age int
  }
  // 给*student实现方法string()
  func (stu *Student) String() string {
      str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name, stu.Age)
      return str
  }
  ```

  ```go
  // 定义一个Student变量
  stu := Student{
      Name : "tom",
      Age : 20,
  }
  //如果实现了 *Student 类型的 String方法，就会自动调用
  fmt.Println(&stu)
  ```


###### 方法和函数的区别：

1. 调用方式不一样

   函数的调用方式：   函数名（实参列表）

   方法的调用方式：   变量.方法名（实参列表）

2. 对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然

3. 对于方法（如struct的方法），接受者为值类型时，可以直接使用指针类型的变量调用方法，反过来同样也可以

```go
func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name)
}
func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test04() =", p.Name)
}

func main() {
	p := Person{"tom"}
	p.test03()
	fmt.Println("main() p.name=", p.Name) //tom
	(&p).test03()                         // 从形式上传入地址，但本质仍然是值拷贝
	fmt.Println("main() p.name=", p.Name) // tom
	(&p).test04()
	fmt.Println("main() p.name=", p.Name) // mary
	p.test04()                            // 等价于(&p).test04()  编译器自动处理，从形式上是传入值类型，但是本质任然是地址拷贝
}
```

总结：

- 不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和那个类型绑定
- 如果是和值类型，比如（p Person），则是值拷贝，如果和指针类型，比如是（p *Person）则是地址拷贝

面向对象编程应用实例：

1. 声明结构体，确定结构体名
2. 编写结构体字段
3. 编写结构体的方法

###### 创建结构体变量时指定字段值：

```go
type Stu struct {
	Name string
	Age  int
}

func main() {

	// 方式1
	// 在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}
	//在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu3 = Stu{
		Name: "jack",
		Age:  20,
	}
	stu4 := Stu{
		Age:  30,
		Name: "mary",
	}
	fmt.Println(stu1, stu2, stu3, stu4)

	//方法2， 返回结构体的指针类型
	var stu5 = &Stu{"小王", 29}
	stu6 := &Stu{"小王~", 39}
	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法不依赖字段的定义顺序
	var stu7 = &Stu{
		Name: "小李",
		Age:  49,
	}
	stu8 := &Stu{
		Age:  59,
		Name: "小李~",
	}
	fmt.Println(*stu5, *stu6, *stu7, *stu8)
}
```

##### 工厂模式：

Go结构体没有构造函数，通常可以使用工厂模式来解决这个问题

```go
package model
type Student struct{
    Name string...
}
```

如果想让类型是小写，而且在其他包里面可以创建这个类型，则需要用到工厂模式

用的时候：

```go
package main
import (
	"fmt"
    "model"
)
func mian() {
    // 定student结构体是首字母小写，我们可以通过工厂模式来解决
    var stu = model.NewStudent("tom~",88)
    fmt.Println(*stu)
    fmt.Println("name=",stu.Name, "score=",stu.Score)
}
```

 model包里：

```go
//因为student结构体首字母是小写，因此是只能在model使用
//我们通过工厂模式来解决
func NewStudent(n string, s float64) *student{
    return &student{
        Name: n,
        Score: s,
        //score: s,
    }
}
// 如果score字段首字母小写，则在其它包不可以直接访问，我们可以提供一个方法
func (s *student) GetScore() float64{
    return s.score
}
```











