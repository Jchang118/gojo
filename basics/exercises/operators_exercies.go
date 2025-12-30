package main

import "fmt"

func main() {
	// 有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
	// 任何数和 00 做异或运算，结果仍然是原来的数，即 a ^ 0=aa⊕0=a。任何数和其自身做异或运算，结果是 00，即 a ^ a=0a⊕a=0。异或运算满足交换律和结合律，即 a ^ b ^ a=b ^ a ^ a=b ^ (a ^ a)=b ^ 0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
	nums := []int{1, 4, 3, 2, 4, 1, 2}
	res := 0
	for _, num := range nums {
		res ^= num
	}
	fmt.Println(res)
}
