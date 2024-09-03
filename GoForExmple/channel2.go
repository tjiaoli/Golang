package main

import (
	"fmt"
	"sync"
)

func main() {
	//有缓冲区的channel，可以通过make的时候指定容量
	ch := make(chan int, 3)
	// 下面2个发送操作不用阻塞等待接收方接收数据,如果是没有缓冲区此处会阻塞
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

	wg.Wait()

	close(ch)
	//关闭后还能取值
	fmt.Println(<-ch) // 20
	fmt.Println(<-ch) // 0
}
