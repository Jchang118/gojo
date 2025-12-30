package main

import "fmt"

func main() {
	// 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型
	a := 5050
	b := 3.1415
	c := true
	d := "Hello World!"

	fmt.Printf("%v (%T)\n", a, a)
	fmt.Printf("%v (%T)\n", b, b)
	fmt.Printf("%v (%T)\n", c, c)
	fmt.Printf("%v (%T)\n", d, d)

	// 编写代码统计出字符串"hello沙河小王子"中汉字的数量
	e := "hello沙河小王子"
	count := 0
	for _, v := range e {
		if len(string(v)) >= 3 {
			count++
		}
	}
	fmt.Println(count)
}
