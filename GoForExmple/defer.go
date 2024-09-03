package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumN(N int) {
	// 调用defer wg.Done()确保sumN执行完之后，可以对wg的计数器减1
	defer wg.Done()
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	fmt.Printf("sum from 1 to %d is %d\n", N, sum)
}

func main() {
	// 设置wg跟踪的计数器数量为1
	wg.Add(1)
	// 开启sumN这个goroutine去计算1到100的和
	//go 函数是并发执行
	go sumN(5)
	//此时输出可能优先于上面的输出
	fmt.Println("finish")
	// Wait会一直等待，直到wg的计数器为0
	wg.Wait()
	fmt.Println("finish")

	i := 0
	defer fmt.Println("i:", i) //虽然defer是最后执行但是被defer的函数的参数在执行defer的时候就被确认了
	i++

	//被defer的函数执行 后进先出（Last In First Out），后进的先执行
	for i := 0; i < 4; i++ {
		if i == 0 {
			defer fmt.Println(i)
		} else {
			defer fmt.Print(i)
		}
	}

	fmt.Println("result:", f())
}

// f returns 42
// 被defer的函数可以对defer语句所在的函数的命名返回值做读取和修改操作
func f() (result int) {
	//defer后面跟的必须是函数或者方法调用，defer后面的表达式不能加括号。
	//defer result *=7
	defer func() {
		result -= 1
	}()
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
