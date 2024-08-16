package main

import "fmt"

func main() {

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("i等于几", i)

	//与java条件省略了()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//range 定义执行几次循环
	for i := range 3 {
		fmt.Println("range 执行", i+1, "次迭代")
	}
	i = 1
	for {
		fmt.Println("loop i = ", i)
		if i == 3 {
			fmt.Println("loop end i =", i, "break 控制 for没有条件反复循环或return")
			break
		}
		if i == 4 {
			fmt.Println("loop end i =", i, "break 控制 for没有条件反复循环或return")
			return
		}
		i++
	}

	for i := range 6 {
		if i%2 == 0 {
			continue //跳过本轮循环
		}
		fmt.Println("小于6的奇数：", i)
	}
}
