package main

import (
	"fmt"
	"time"
)

// 单向通道
// chan <- int // 只写，只能往channel写数据，不能从channel读数据
// <- chan int // 只读，只能从channel读数据，不能往channel写数据
func write(ch chan<- int) {
	ch <- 10
}

func read(ch <-chan int) {
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan int)
	go write(ch)
	//read(ch)

	fmt.Println(<-ch)

	time.Sleep(time.Second)
}
