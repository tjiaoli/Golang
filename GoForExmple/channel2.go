package main

import (
	"fmt"
	"sync"
)

func main() {
	//有缓冲区的channel，可以通过make的时候指定容量
	ch := make(chan int, 2)
	// 下面2个发送操作不用阻塞等待接收方接收数据
	ch <- 10
	ch <- 20
	/*
		如果添加下面这行代码，就会一直阻塞，因为缓冲区已满，运行会报错
		fatal error: all goroutines are asleep - deadlock!

		ch <- 30
	*/

	var wg sync.WaitGroup

	fmt.Println(<-ch) // 10

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 30
		value := <-ch
		fmt.Println("value:", value)
	}()

	fmt.Println(<-ch) // 20

	wg.Wait()
	ch <- 30

	ch <- 4
}
