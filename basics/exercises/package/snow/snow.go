package snow

import (
	"fmt"
	"package/calc"
)

func Snow() {
	var x, y int
	fmt.Printf("请输入第一个整数: ")
	fmt.Scanf("%d\n", &x)
	fmt.Printf("请输入第二个整数: ")
	fmt.Scanf("%d\n", &y)
	fmt.Printf("两数相加结果为: %d\n", calc.Add(x, y))
	fmt.Printf("两数相减结果为: %d\n", calc.Sub(x, y))
	fmt.Printf("两数相乘结果为: %d\n", calc.Mul(x, y))
	fmt.Printf("两数相除结果为: %d\n", calc.Div(x, y))
}
