package main

import "fmt"

// Go 内置了对多个返回值的支持。此功能在惯用 Go 中经常使用，例如从函数返回结果和错误值。

// 该(int, int)函数签名表明该函数返回 2 ints。
func vals() (int, int) {
	return 3, 7
}

func main() {
	//这里我们使用了来自多重赋值的调用的两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果您只想要返回值的子集，请使用空白标识符_。
	_, c := vals()
	fmt.Println(c)
}
