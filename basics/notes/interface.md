**Go语言基础之接口**
接口(interface)定义了一个对象的行为规范,只定义规范不实现,由具体的对象来实现规范的细节.

# 接口
本章学习目标
- 了解为什么需要接口以及接口的特点
- 掌握接口的声明和使用
- 掌握接口值的概念
- 掌握空接口的特点以及其使用场景

在Go语言中接口(interface)是一种类型,一种抽象的类型.相较于之前章节中讲到的那些具体类型(字符串、切片、结构体等)更注重“我是谁”,接口类型更注重“我能做什么”的问题.接口类型就像是一种约定--概括了一种类型应该具备哪些方法,在Go语言中提倡使用面向接口的编程方式实现解耦.

## 接口类型
接口是一种由程序员来定义的类型,一个接口类型就是一组方法的集合,它规定了需要实现的所有方法.

相较于使用结构体类型,当我们使用接口类型说明相比于它是什么更关心它能做什么.

### 接口的定义
每个接口类型由任意个方法签名组成,接口的定义格式如下:
```go
type 接口类型名 interface {
    方法名1( 参数列表 ) 返回值列表1
    方法名2( 参数列表 ) 返回值列表2
    ...
}
```
其中:
- 接口类型名: Go语言的接口在命名时,一般会在单词后面添加er,如有写操作的接口叫Writer,有关操作的接口叫closer.接口名最好要能突出该接口的类型含义.
- 方法名: 当方法名首字母是大写且这个接口类型名首字母也是大写时,这个方法可以被接口所在的包(package)之外的代码访问.
- 参数列表、返回值列表: 参数列表和返回值列表中的参数变量名可以省略.

举个例子,定义一个包含`Write`方法的`Writer`接口.
```go
type Writer interface {
    Write([]byte) error
}
```
当你看到一个`Writer`接口类型的值时,你不知道它是什么,唯一知道的就是可以通过调用它的`Write`方法来做一些事情.

### 实现接口的条件
接口就是规定了一个需要实现的方法列表,在Go语言中一个类型只要实现了接口中规定的所有方法,那么我们就称它实现了这个接口.

我们定义的`Singer`接口类型,它包含了一个`Sing`方法.
```go
// Singer 接口
type Singer interface {
    Sing()
}
```
我们有一个`Bird`结构体类型如下.
```go
type Bird struct {}
```
因为`Singer`接口只包含了一个`Sing`方法,所以只需要给`Bird`结构体添加一个`Sing`方法就可以满足`Singer`接口的要求.
```go
// Sing Bird类型的Sing方法
func (b Bird) Sing() {
    fmt.Println("汪汪汪")
}
```
这样就称为`Bird`实现了`Singer`接口.

### 为什么要使用接口
现在假设我们的代码世界里有很多小动物,下面的代码片段定义了猫和狗,它们饿了都会叫.
```go
package main

import "fmt"

type Cat struct{}

func (c Cat) Say() {
    fmt.Println("喵喵喵")
}

type Dog struct{}

func (d Dog) Say() {
    fmt.Println("汪汪汪")
}

func main() {
    c := Cat{}
    c.Say()
    d := Dog{}
    d.Say()
}
```
这个时候又跑来了一只羊,羊饿了也会发出叫声.
```go
type Sheep struct{}

func (s Sheep) Say() {
    fmt.println("咩咩咩")
}
```
我们接下来定义一个饿肚子的场景.
```go
// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
    c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
    s.Say()
}
```
接下来会有越来越多的小动物跑过来,我们的代码世界该怎么拓展呢?

在饿肚子这个场景下,我们可不可以把所有动物都当成一个“会叫的类型“来处理呢?当然可以!使用接口类型就可以实现这个目标.我们的代码其实并不关心究竟是什么动物在叫,我们只是在代码中调用它的`Say()`方法,这就足够了.

我们可以约定一个`Sayer`类型,它必须实现一个`Say()`方法,只要饿肚子了,我们就调用`Say()`方法.
```go
type Sayer interface {
    Say()
}
```
然后我们定义一个通用的`MakeHungry`函数,接收`Sayer`类型的参数.
```go
// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
    s.Say()
}
```
