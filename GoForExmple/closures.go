package main

import "fmt"

// 闭包函数 func x() func() type||null { i := type-val return func() type||null { }}
// 这个函数 intSeq 返回另一个函数，该函数在 intSeq 的函数体内以匿名方式定义。返回的这个函数会闭包引用变量 i，从而形成一个闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//我们调用intSeq，将结果（一个函数）赋值给nextInt。此函数值捕获其自身的i值，每次调用 时该值都会更新nextInt。
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(nextInt())

	outer := outer()
	//无返回类型
	outer()

	outer()
}

func stringSeq() func() string {
	s := "Hello"
	return func() string {
		s += "!"
		return s
	}
}

func outer() func() {
	capturedVariable := 0 // 这是在外层函数中定义的变量
	return func() {
		// 这个匿名函数形成一个闭包，捕获了 capturedVariable
		capturedVariable++
		fmt.Println(capturedVariable)
	}
}
