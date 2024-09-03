package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		x := <-ch // 从通道ch里接收值，并赋值给变量x
		fmt.Println(x)
	}()

	//ch := make(chan int)
	ch <- 10 // 把10发送到ch里

	a := 10
	ch1 := make(chan int)
	go func() {
		defer wg.Done()
		ch1 <- a
	}()
	value := <-ch1

	fmt.Println("value:", value)

	wg.Wait()
}
