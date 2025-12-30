package main

import "fmt"

func main() {
	// 求数组[1, 3, 5, 7, 8]所有元素的和
	a := [...]int{1, 3, 5, 7, 8}
	count := 0
	for _, v := range a {
		count += v
	}
	fmt.Println(count)

	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
