package main

import "time"

func addData(ch chan int) {
	/*
		每3秒往通道ch里发送一次数据
		cap(chan) channel的容量
	*/
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
	// 数据发送完毕，关闭通道
	close(ch)
}

//func main() {
//	ch := make(chan int, 10)
//	// 开启一个goroutine，用于往通道ch里发送数据
//	go addData(ch)
//
//	/* range迭代从通道ch里获取数据
//	通道close后，range迭代取完通道里的值后，循环会自动结束
//	如果通道未关闭，‘range’循环将永远不会结束，因为它会一直等待新数据
//	*/
//	for i := range ch {
//		fmt.Println(i)
//	}
//}
