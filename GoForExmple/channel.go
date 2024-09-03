package main

import "fmt"
import "time"

type Cat struct {
	name string
	age  int
}

func fetchChannel(ch chan Cat) {
	value := <-ch
	fmt.Printf("type: %T, value: %v\n", value, value)
}

func fetchChannel1(a Cat, ch chan Cat) {
	ch <- a
}

func main() {

	ch1 := make(chan int)
	ch1 <- 10 // 把10发送到ch里
	fmt.Println("000")
	ch := make(chan Cat)
	a := Cat{"yingduan", 1}
	fetchChannel1(a, ch)
	// 启动一个goroutine，用于从ch这个通道里获取数据
	go fetchChannel(ch) //通道操作会阻塞线程，goroutine并发value:= <-ch 接收通道的数据，需要等待ch <-a 给通道发送数据

	fmt.Println(a)
	fmt.Println("11111")
	// 往cha这个通道里发送数据，ch接收数据后fetchChannel继续往下执行打印
	//ch <- a
	// main这个goroutine在这里等待2秒
	time.Sleep(2 * time.Second)
	fmt.Println("end")

}
