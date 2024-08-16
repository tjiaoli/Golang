package main

import "fmt"

func main() {
	var a int = 100
	fmt.Println(a)
	var b int = 20000000000
	fmt.Println(a + b)
	var c = "inital"
	fmt.Println(c)
	var d = true
	fmt.Println(d)
	var e = false
	fmt.Println(e)

	//声明和初始化的简写。此语法仅在函数内部可用
	f := "inital"
	fmt.Println(f)
	f = "11112"
	fmt.Println(f)
	var g int
	fmt.Println(g)
	//g := 9 已经声明的默认初始化为0
	//for 循环中的g与外面的g不冲突
	for g := 7; g < 10; g++ {

	}
	g = 9
	fmt.Println(g)

	h := 10
	fmt.Println(h)
}
