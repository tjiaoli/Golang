package main

import "fmt"

// 这是一个将任意数量的ints 作为参数的函数。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	//在函数内部， 的类型nums相当于[]int。我们可以调用len(nums)，用 对其进行迭代range，等等。
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func stri(str ...string) {
	fmt.Print(str, " ")

	//for _, str := range str {
	//	fmt.Println(str, " ")
	//}
}
func main() {

	//可变参数函数可以按照通常的方式使用单独的参数来调用。
	sum(1, 2)
	sum(1, 2, 3)

	//如果切片中已经有多个参数，请使用 func(slice...)如下方式将它们应用于可变函数。
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	stri("23", "33")

	strs := []string{"22", "33", "44"}

	stri(strs...)
}
