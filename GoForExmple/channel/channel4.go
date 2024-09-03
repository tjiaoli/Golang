package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	// 开启一个goroutine，用于往通道ch里发送数据
	go addData(ch)

	/*
		for循环取完channel里的值后，因为通道close了，再次获取会拿到对应数据类型的零值
		如果通道不close，for循环取完数据后就会阻塞报错
	*/
	for {
		//ok是波尔值用于判断通道是否关闭
		value, ok := <-ch
		if ok {
			fmt.Println(ok)
			fmt.Println(value)
		} else {
			fmt.Println("finish")
			break
		}
	}

}
