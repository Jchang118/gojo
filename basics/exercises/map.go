package main

import (
	"fmt"
	"strings"
)

func main() {
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1
	wordMap := make(map[string]int)
	s1 := "how do you do"
	wordSlice := strings.Fields(s1)
	for _, word := range wordSlice {
		_, ok := wordMap[word]
		if ok {
			wordMap[word]++
		} else {
			wordMap[word] = 1
		}
	}
	for word, count := range wordMap {
		fmt.Printf("%s=%d ", word, count)
	}
	fmt.Println()

	// 观察下面代码，写出最终的打印结果
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s               // 这里拷贝了slice header(三个字段:指针ptr, len, cap),所以m["q1mi"]的len还是3
	s = append(s[:1], s[2:]...) //s的len变成了2,cap还是3
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
	// [1 2 3]
	// [1 3]
	// [1 3 3]
}
