package main

import "fmt"

func plus(a int, b int) int {
	c := a + b
	return c
}

// 当您有多个连续的相同类型的参数时，您可以省略同类型参数的类型名称，直到声明类型的最后一个参数
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
