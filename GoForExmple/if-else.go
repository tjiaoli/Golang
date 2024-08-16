package main

import (
	"fmt"
)

func main() {

	//Go 中没有三元 ifif ，因此即使是基本条件也 需要使用完整语句。

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4") //8能被4整除
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//语句可以位于条件之前；此语句中声明的任何变量均可在当前分支和所有后续分支中使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	num := -1
	//可以声明多个变量但是必须是同一种类型，或能推到成同一种类型
	if v, v1 := "2", "33"; num < 0 {
		fmt.Println(v, v1)
	}

	var a int = 5
	var b float64 = 10.5

	if floatA, floatB := float64(a), 11.3; floatA > b {
		fmt.Println(a, "is too small")
	} else if floatB > b {
		fmt.Println(floatB, "is too big")
	}
}
