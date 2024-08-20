package main

import "fmt"

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(7))

	//闭包也可以是递归的，但这要求在定义闭包之前，使用带有类型的 `var` 明确声明该闭包。
	var fib func(n int) int
	//斐波那契数列，兔子繁殖 f(n) = f(n-1) + f(n-2) 效率较低 ，需要先算出 f(0) +f(1) 然后再回溯相加
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	//通过这种方式，时间复杂度从 O(2^n) 降低到 O(n)。
	var memo = map[int]int{}

	var fib1 func(n int) int
	fib1 = func(n int) int {
		if n < 2 {
			return n
		}
		if result, exists := memo[n]; exists {
			return result
		}
		memo[n] = fib1(n-1) + fib1(n-2)
		return memo[n]
	}

	fmt.Println(fib1(7))

	globalFib := fib
	fmt.Println(globalFib(7))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	prev1, prev2 := 0, 1
	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev1 = prev2
		prev2 = current
	}
	return prev2
}
