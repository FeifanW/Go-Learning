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

4. 结构体嵌入两个（或多个）匿名结构体，如果两个匿名结构体有相同的字段和方法（同时结构体半身没有同名的字段和方法）,在访问时，就必须明确指明匿名结构体名字否则编译报错

   ```go
   type A struct {
       Name string
       Age int
   }
   type B struct {
       Name string
       score int
   }
   type C struct {
       A
       B
       // Name string
   }
   ```

5. 如果一个struct嵌套了一个有名结构体，这种模式就是**组合**，如果是组合关系，哪么在访问组合的结构体的字段或方法时，必须带上结构体的名字

   ```go
   type A struct {
       Name string
       Age int
   }
   type C struct {
       a A
   }
   ```

6. 嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值

   ```go
   type Goods struct {
   	Name  string
   	Price float64
   }
   
   type Brand struct {
   	Name    string
   	Address string
   }
   
   type TV struct {
   	Goods
   	Brand
   }
   
   type TV2 struct {
   	*Goods
   	*Brand
   }
   
   func main() {
   	// 嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
   	tv := TV{Goods{"电视机", 5000.99}, Brand{"海尔", "山东"}}
   	tv2 := TV{
   		Goods{
   			Name:  "电视机",
   			Price: 5000.99,
   		},
   		Brand{
   			Name:    "海尔",
   			Address: "山东",
   		},
   	}
   	fmt.Println("tv", tv)
   	fmt.Println("tv2", tv2)
   
   	tv3 := TV2{&Goods{"电视机", 7000.5}, &Brand{"创维", "河南"}}
   	tv4 := TV2{
   		&Goods{
   			Name:  "电视机",
   			Price: 7000.5,
   		},
   		&Brand{"创维", "河南"}}
   	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
   	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
   }
   ```

   ```
   type Monster struct {
   	Name string
   	Age  int
   }
   
   type E struct {
   	Monster
   	int
   }
   
   func main() {
   	var e E
   	e.Name = "狐狸"
   	e.Age = 300
   	e.int = 20
   	e.n = 40
   	fmt.Println("e=",e)
   }
   ```
   
   说明：
   
   1. 如果一个结构体有int类型的匿名字段，就不能第二个
   2. 如果需要有多个int字段，则必须给int字段
   
   多重继承说明：
   
   如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承

###### 接口：

按照顺序应该讲解多态，但是讲解多态前，需要讲解接口，因为Go中多态特性主要是通过接口体现的

```go
package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
}

// 让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

// 让Camera实现 Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 计算机
type Computer struct {
}

// 编写一个方法Working方法，接收一个Usb接口类型变量
// 只要实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明的所有方法）
func (c Computer) Working(usb Usb) {
	// 通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}
func main() {
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
```

interface类型可以定义一组方法，但是这些不需要实现，并且interface不能包含任何变量，到某个自定义类型要使用的时候，在根据具体情况把这些方法写出来（实现）

基本语法：

```go
type 接口名 interface {
    method1(参数列表) 返回值列表
    method2(参数列表) 返回值列表
    ...
}
```

实现接口所有方法

```go
func (t 自定义类型) method1(参数列表) 返回值列表 {
    // 方法实现
}
func (t 自定义类型) method2(参数列表) 返回值列表 {
    // 方法实现
}
// ...
```

小结说明：

1. 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法，接口体现了程序设计的**多态和高内聚低耦合**的思想
2. Go中的接口，不需要显示的实现，只要一个变量，含义接口类型中的所有方法，那么这个变量就实现这个接口，因此**Go中没有implement这样的关键字**

注意细节：

1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）

2. 接口中所有的方法都没有方法体，即都是没有实现的方法

3. 在Go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口

4. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型

5. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

6. 一个自定义类型可以实现多个接口

   ```go
   type AInterface interface {
       Say()
   }
   type BInterface interface {
       Hello()
   }
   type Monster struct {
       
   }
   func (m Monster) Hello() {
       fmt.Println("Monster Hello()")
   }
   func (m Monster) Say() {
       fmt.Println("Monster Say()")
   }
   // Monster实现了AInterface和BInterface
   var monster Monster
   var a2 AInterface = monster
   var b2 BInterface = monster
   a2.Say()
   b2.Hello()
   ```

7. Go接口中不能有任何变量

8. 一个接口（比如A接口）可以继承多个别的接口（比如B,C接口），这是如果要实现A接口，也必须将B，C接口的方法也全部实现

9. interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，那么会输出nil

10. 空节课interface{}没有任何方法，所以所有类型都实现了空接口，即我们可以把任何变量赋给空接口

    ```go
    type Usb interface {
    	Say()
    }
    type Stu struct {
    }
    
    func (this *Stu) Say() {
    	fmt.Println("Say()")
    }
    func main() {
    	var stu Usb = &Stu{}
    	//var stu Stu = Stu{}
    	// 错误！会报Stu类型没有实现Usb接口
    	// 如果希望通过编译 var u Usb = &stu
    	var u Usb = stu
    	u.Say()
    	fmt.Println("here", u)
    }
    ```

    ```go
    package main
    
    import (
    	"fmt"
    	"math/rand"
    	"sort"
    )
    
    // 1.声明Hero结构体
    type Hero struct {
    	Name string
    	Age  int
    }
    
    // 2.声明一个Hero结构体切片类型
    type HeroSlice []Hero
    
    // 3.实现Interface接口
    func (hs HeroSlice) Len() int {
    	return len(hs)
    }
    
    //Less方法就是决定你使用什么标准进行排序
    //1.按照Hero的年龄从小到大排序
    func (hs HeroSlice) Less(i, j int) bool {
    	return hs[i].Age > hs[j].Age
    }
    
    func (hs HeroSlice) Swap(i, j int) {
    	//temp := hs[i]
    	//hs[i] = hs[j]
    	//hs[j] = temp
    	hs[i], hs[j] = hs[j], hs[i]
    }
    func main() {
    	// 先定义一个数组/切片
    	var intSlice = []int{0, -1, 10, 7, 90}
    	// 要求对intSlice切片进行排序
    	// 1.冒泡排序
    	// 2.也可以使用系统提供的方法
    	sort.Ints(intSlice)
    	fmt.Println(intSlice)
    
    	// 测试看看我们是否可以对结构体切片进行排序
    	var heroes HeroSlice
    	for i := 0; i < 10; i++ {
    		hero := Hero{
    			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
    			Age:  rand.Intn(100),
    		}
    		//将hero append到heroes切片
    		heroes = append(heroes, hero)
    	}
    	// 看看排序前的顺序
    	for _, v := range heroes {
    		fmt.Println(v)
    	}
    	//调用sort.Sort
    	sort.Sort(heroes)
    	// 看看排序后的顺序
    	for _, v := range heroes {
    		fmt.Println(v)
    	}
    }
    ```

    























