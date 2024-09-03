package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func main() {
	/*开启一个goroutine去执行hello函数*/
	go hello()
	/*go会为main函数创建一个默认的goroutine，如果main()函数结束了，那所有的main()中的goroutine都会立即结束*/
	fmt.Println("main end")
}
