package main

import (
	"fmt"
	"sort"
)

func main() {
	// 以下代码块输出为["" "" "" "" "" 0 1 2 3 4 5 6 7 8 9]
	// 因为make([]string, 5, 10)已经把slice长度设成5了,append是在“后面继续加”
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)

	// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	var b = [...]int{3, 7, 8, 9, 1}
	s := b[:]
	sort.Ints(s)
	fmt.Println(b)
}
