# Go语言基础之并发
并发编程在当前软件领域是一个非常重要的概念,随着CPU等硬件的发展,我们无一例外的想让我们的程序运行的快一点,再快一点.Go语言在语言层面天生支持并发,充分利用现代CPU的多核优势,这也是Go语言能够大范围流行的一个很重要的原来.

## 基本概念
首先我们先来了解介个与并发编程相关的基本概念.

### 串行,并发与并行
串行: 我们都是先读小学,小学毕业后再读初中,读完初中再读高中.
并发: 同一时间段内执行多个任务(你在用微信和两个女朋友聊天).
并行: 同一时刻执行多个任务(你和你朋友都在用微信和女朋友聊天).

### 进程,线程和协程
进程(process): 程序在操作系统中的一次执行过程,系统进行资源分配和调度的一个独立单位.
线程(thread): 操作系统基于进程开启的轻量级进程,是操作系统调度执行的最小单位.
协程(coroutine): 非操作系统提供而是由用户自行创建和控制的用户态'线程',比线程更轻量级.

### 并发模型
业界将如何实现并发编程总结归纳为各式各样的并发模型,常见的并发模型有以下几种:
- 线程&锁模型
- Actor模型
- CSP模型
- Fork&Join模型

Go语言中的并发程序主要是通过基于CSP(communicating sequential process)的goroutine和channel来实现,当然也支持使用传统的多线程共享内存的并发方式.

## goroutine
Goroutine是Go语言支持并发的核心,在一个Go程序中同时创建成百上千个goroutine是非常普遍的,一个goroutine会以一个很小的栈开始其生命周期,一般只需要2KB.区别于操作系统线程由系统内核进行调度,goroutine是由Go运行时(runtime)负责调度.例如Go运行时会智能地将m个goroutine合理地分配给n个操作系统线程,实现类似m:n地调度机制,不再需要Go开发者自行在代码层面维护一个线程池.

Goroutine是Go程序中最基本地并发执行单元.每一个Go程序都至少包含一个gorountine--main goroutine,当Go程序启动时它会自动创建.

在Go语言编程中你不需要去自己写进程,线程,协程,你的技能包里只有一个技能--goroutine,当你需要让某个任务并发执行地时候,你只需要把这个任务包装成一个函数,开启一个goroutine去执行这个函数就可以了,就是这么简单粗暴.

### go关键字
Go语言中使用goroutine非常简单,只需要在函数或方法调用前加上`go`关键字就可以创建一个goroutine,从而让该函数或方法在新创建的goroutine中执行.
```go
go f()  // 创建一个新的goroutine去执行.
```
匿名函数也支持使用`go`关键字创建goroutine去执行.
```go
go func(){
    // ...
}()
```
一个goroutine必定对应一个函数/方法,可以创建多个goroutine去执行相同的函数/方法.

### 启动单个goroutine
启动单个goroutine的方式非常简单,只需要在调用函数(普通函数和匿名函数)前加上一个`go`关键字.

我们先来看一个在main函数中执行普通函数调用的示例.
```go
package main

import "fmt"

func hello() {
    fmt.Println("hello")
}

func main() {
    hello()
    fmt.Println("你好")
}
```
将上面的代码编译后执行,得到的结果如下:
```
hello
你好
```
代码中hello函数和其后面的打印语句是串行的.
![goroutine01](img/goroutine01.png)

接下来我们在调用hello函数前面加上关键字`go`,也就是启动一个goroutine去执行hello这个函数.
```go
func main() {
    go hello()  // 启动另外一个goroutine去执行hello函数
    fmt.Prtinln("你好")
}
```
将上述代码重新编译后执行,得到输出结果如下:
```
你好
```
这一次的执行结果只在终端打印了“你好”,并没有打印`hello`.这是为什么呢?

其实在Go程序启动时,Go程序就会为main函数创建一个默认的goroutine.在上面的代码中我们在main函数中使用go关键字创建了另外一个goroutine去执行hello函数,而此时main goroutine还在继续往下执行,我们的程序中此时存在两个并发执行的goroutine.当main函数结束时整个程序也就结束了,同时main goroutine也结束了,所有由main goroutine创建的goroutine也会一同退出.也就是说我们的main函数退出太快,另外一个goroutine中的函数还未执行完程序就退出了,导致未打印出“hello”.

*main goroutine就像是《权利的游戏》中的夜王,其他的goroutine都是夜王转化出的异鬼,夜王一死它转化的那些异鬼也就全部GG了.*

所以我们要想办法让main函数“等一等”将在另一个`goroutine`中运行的`hello`函数.其中最简单粗暴的方式就是在main函数中“time.Sleep”1秒钟(这里的1秒钟是我们根据经验而设置的一个值,在这个示例中1秒钟足够创建新的`goroutine`执行完`hello`函数了).

按如下方式修改我们的示例代码:
```go
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("hello")
}

func main() {
    go hello()
    fmt.Println("你好")
    time.Sleep(time.Second)
}
```
将我们的程序重新编译后再次执行,程序会在终端输出如下结果,并且会短暂停顿一会儿.
```
你好
hello
```
为什么会先打印`你好`呢?

这是因为在程序中创建goroutine执行函数需要一定的开销,而与此同时main函数所在的goroutine是继续执行的.
![goroutine02](img/goroutine02.png)

在上面的程序中使用`time.Sleep`让main goroutine等待hello goroutine执行结束是不优雅的,当然也不是准确的.

Goy语言中通过`sync`包为我们提供了一些常用的并发原语,我们会在后面的小节单独介绍`sync`包中的内容.在这一小节,我们会先介绍一下sync包中的`WaitGroup`.当你并不关心并发操作的结果或者有其他方式收集并发操作的结果时,`WaitGroup`是实现等待一组并发操作完成的好方法.

下面的示例代码中我们在main goroutine中使用`sync.WaitGroup`来等待hello routine完成后再退出.
```go
package main

import (
    "fmt"
    "sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
    fmt.Println("hello")
    wg.Done()   // 告知当前goroutine完成
}

func main() {
    wg.Add(1)   // 登记1个goroutine
    go hello()
    fmt.Println("你好")
    wg.Wait()   // 阻塞等待登机的goroutine完成
}
```
将代码编译后再执行,得到的输出结果和之前一致,但是这一次程序不再会有多余的停顿,hello goroutine执行完毕后程序直接退出.

### 启动多个goroutine
在Go语言中实现并发就是这样简单,我们还可以启动多个goroutine.让我们再来看一个新的代码示例.这里同样使用了`sync.WaitGroup`来实现goroutine的同步.
```go
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func hello(i int) {
    defer wg.Done() // goroutine结束就登记-1
    fmt.Println("hello", i)
}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1)   // 启动一个goroutine就登记+1
        go hello(i)
    }
    wg.Wait()   // 等待所有登记的goroutine都结束
}
```
多次执行上面的代码会发现每次终端上打印数字的顺序都不一致.这是因为10个goroutine是并发执行的,而goroutine的调度是随机的.

### 动态栈
操作系统的线程一般都有固定的栈内存(通常为2MB),而Go语言中的goroutine非常轻量级,一个goroutine的初始栈空间很小(一般为2KB),所以在Go语言中一次创建数万个goroutine也是可能的.并且goroutine的栈不是固定的,可以根据需要动态地增大或缩小,Go的runtime会自动为goroutine分配合适的栈空间.

### goroutine调度
操作系统内核在调度时会挂起当前正在执行的线程并将寄存器中的内容保存到内存中,然后选出接下来要执行的线程并从内存中恢复该线程的寄存器信息,然后恢复执行该线程的现场并开始执行线程.从一个线程切换到另一个线程需要完成的上下文切换.因此可能需要多次内存访问,索引这个切换上下文的操作开销比较大,会增加运行的cpu周期.

区别于操作系统内核调度操作系统线程,goroutine的调度是Go语言运行时(runtime)层面的实现,是完全由Go语言本身实现的一套调度系统--go scheduler.它的作用是按照一定的规则将所有的goroutine调度到操作系统线程上执行.

在经历数个版本的迭代之后,目前Go语言的调度器采用的是`GPM`调度模型.
![gpm](img/gpm.png)
其中:
- G: 代表goroutine,每执行一次`go f()`就创建一个G,包含要执行的函数和上下文信息.
- 全局队列(Global Queue): 存放等待运行的G.
- P: 表示goroutine执行所需的资源,最多有GOMAXPROCS个.
- P的本地队列: 同全局队列类似,存放的也是等待运行的G,存的数量有限,不超过256个.新建G时,G优先加入到P的本地队列,如果本地队列满了会批量移动部分G到全局队列.
- M: 线程想运行任务就获得P,从P的本地队列获取G,当P的本地队列为空时,M也会尝试从全局队列或其他P的本地队列获取G.M运行G,G执行之后,M会从P获取下一个G,不断重复下去.
- Goroutine调度器和操作系统调度器时通过M结合起来的,每个M都代表了1个内核线程,操作系统调度器负责把内核线程分配到CPU的核上执行.

单从线程调度讲,Go语言相比其他语言的优势在于OS线程是由OS内核来调度的,goroutine则是Go运行时(runtime)自己的调度器调度的,完全是在用户态下完成的,不涉及内核态于用户态之间的频繁切换,包括内存的分配与释放,都是在用户态维护着一块大的内存池,不直接调用系统的malloc函数(除非内存池需要改变),成本比调度OS线程低很多.另一方面充分利用了多核的硬件资源,近似的把若干goroutine均分在物理线程上,再加上本身goroutine的超轻量级,以上种种特性保证了goroutine调度方面的性能.

### GOMAXPROCS
Go运行时的调度器使用`GOMAXPROCS`参数来确定需要使用多少个OS线程来同时执行Go代码.默认值是机器上的CPU核心数.例如在一个8核心的机器上,GOMAXPROCS默认为8.Go语言中可以通过`runtime.GOMAXPROCS`函数设置当前程序并发时占用的CPU逻辑核心数.(Go1.5版本之前,默认使用的是单恶心执行.Go1.5版本之后,默认使用全部的CPU逻辑核心数.)

## channel
单纯地将函数并发执行是没有意义的.函数与函数间需要交换数据才能体现并发执行函数的意义.

虽然可以使用共享内存进行数据交换,但是共享内存在不同的goroutine中容易发生竞态问题.为了保证数据交换的正确性,很多并发模型中必须使用互斥量对内存进行加锁,这种做法势必造成性能问题.

Go语言采用的并发模型是`CSP(Comunicating Sequential Processes`,提倡通过通信共享内存而不是通过共享内存而实现通信.

如果说goroutine是Go程序并发的执行体,`channel`就是它们之间的连接.`channel`是可以让一个goroutine发送特定值到另一个goroutine的通信机制.

Go语言中的通道(channel)是一种特殊的类型.通道像一个传送带或者队列,总是遵循先入先出(First In First Out)的规则,保证收发数据的顺序.每一个通道都是一个具体类型的导管,也就是声明channel的时候需要为其指定元素类型.

### channel类型
`channel`是Go语言中一种特有的类型.声明通道类型变量的格式如下:
```go
var 变量名称 chan 元素类型
```
其中:
- chan: 是关键字
- 元素类型: 是指通道中传递元素的类型

举几个例子:
```go
var ch1 chan int    // 声明一个传递整型的通道
var ch2 chan bool   // 声明一个传递布尔型的通道
var ch3 chan []int  // 声明一个传递int切片的通道
```
### channel零值
未初始化的通道类型变量其默认零值是`nil`.
```go
var ch chan int
fmt.Println(ch) // <nil>
```
### 初始化channel
声明的通道类型变量需要使用内置的`make`函数初始化之后才能使用.具体格式如下:
```go
make(chan 元素类型, [缓冲大小])
```
其中:
- channel的缓冲大小是可选的.

举几个例子:
```go
ch4 := make(chan int)
ch5 := make(chan bool, 1)   // 声明一个缓冲区大小为1的通道
```
### channel操作
通道共有发送(send)、接收(receive)和关闭(close)三种操作.而发送和接收操作都使用`<-`符号.

现在我们先使用一下语句定义一个通道:
```go
ch := make(chan int)
```
**发送**
将一个值发送到通道中.
```go
ch <- 10    // 把10发送到ch中
```
**接收**
从一个通道中接收值.
```go
x := <- ch  // 从ch中接收值并赋值给变量x
<-ch        // 从ch中接收值,忽略结果
```
**关闭**
我们通过调用内置的`close`函数来关闭通道.
```go
close(ch)
```
**注意:** 一个通道值是可以背垃圾回收掉的.通道通常由发送方执行关闭操作,并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作.它和关闭文件不一样,通常在结束操作之后关闭文件是必须要做的,但关闭通道不是必须的.

关闭后的通道有以下特点:
- 对一个关闭的通道再发送值就会导致panic.
- 对一个关闭的通道进行接收会一直获取值直到通道为空.
- 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值.
- 关闭一个已经关闭的通道会导致panic.

### 无缓冲的通道
无缓冲的通道又
