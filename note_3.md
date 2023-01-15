#### 一、面向对象：

抽象：把一类事物的共有属性（字段）和行为（方法）抽取出来，形成一个物理模型（模板），这种研究问题的方法称为抽象

##### 1、面向对象的三大特性：

继承、封装和多态

###### 封装：

就是把抽象出的字段和对字段的操作封装在一起，程序被保护在内部，程序的其他包只能通过被授权的操作（方法），才能对字段进行操作

封装的好处：

1. 隐藏实现细节
2. 可以对数据就进行验证，保证安全合理

如何体现封装：

1. 对结构体中的属性进行封装
2. 通过方法，包实现封装

封装的实现步骤：

1. 将结构体、字段（属性）的首字母小写（不能导出了，其它包不能使用，类似private）

2. 给结构体所在包提供一个工厂模式的函数，首字母大写，类似一个构造函数

3. 提供一个首字母大写的Set方法（类似其他语言的public），用于对属性判断并赋值

   ```go
   func (var 结构体类名) SetXxx(参数列表)(返回值列表){
       //加入数据验证的业务逻辑
       var.字段 = 参数
   }
   ```

4. 提供一个首字母大写的Get方法（类似其他语言的public），用于获取属性的值

   ```go
   func (var 结构体类型名) GetXxx(){
       return var.age
   }
   ```

案例：

新建encapsulate/main/main.go和encapsulate/model/person.go两个文件

main.go里面：

```go
package main

import (
	"Go-Learning/encapsulate/model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
}
```

person.go里面：

```go
package model

import "fmt"

type person struct {
	Name string
	age  int // 其他包不能直接访问
	sal  float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和sal我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确...")
	}
}
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确...")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
```

###### 继承：

继承可以解决代码复用，当多个结构体存在相同的属性（字段）和方法时，可以抽象出结构体，在结构体中定义这些相同的属性和方法

在Go中，如果一个struct嵌套了另一个匿名结构体，哪么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性

```go
type Goods struct{
    Name string
    Price int
}
type Book struct {
    Goods  // 这里就是嵌套匿名结构体Goods
	Writer string
}
```

便利：

1. 代码的复用性提高了
2. 代码的扩展性和维护性提高了

继承的深入讨论：

1. 结构体可以使用嵌套匿名结构体的所有字段和方法，即首字母大写或者小写的字段、方法都可以使用

   ```go
   package main
   
   import "fmt"
   
   type A struct {
   	Name string
   	age  int
   }
   
   func (a *A) SayOk() {
   	fmt.Println("A SayOk", a.Name)
   }
   
   func (a *A) hello() {
   	fmt.Println("A hello", a.Name)
   }
   
   type B struct {
   	A
   }
   
   func main() {
   	var b B
   	b.A.Name = "tom"
   	b.A.age = 19
   	b.A.SayOk()
   	b.A.hello()
   }
   
   ```

2. 匿名结构体字段访问可以简化

   b.A.name = "tom" => b.name = "tom"

3. 当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则，如希望访问匿名结构体的字段和方法，可以通过匿名结构体来区分























