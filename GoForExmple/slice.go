package main

import (
	"fmt"
	"unsafe"
)

// slice 切片底层是结构体，里面实际有个指针array，类型是unsafe.Pointer，也就是个指针，指向存放数据的数组。
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func printSlice(param []int) {
	fmt.Printf("slice len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
}

func main() {
	slice1 := []int{1}
	slice2 := make([]int, 3, 100)
	slice2 = []int{1, 2, 3, 4, 5, 6}
	slice1 = append(slice1, slice2...)
	printSlice(slice1)
	printSlice(slice2)

	slice := make([]int, 3, 100)
	var slice3 []int

	fmt.Println("slice==nil", slice == nil) // false
	printSlice(slice)

	fmt.Println("slice3==nil", slice3 == nil) // true
	printSlice(slice3)

	/*下标访问切片*/
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}

	/*range迭代访问切片*/
	for index, value := range slice {
		fmt.Printf("slice[%d]=%d\n", index, value)
	}

	a := make([]int, 0, 4)
	b := append(a, 1)    // b=[1], a指向的底层数组的首元素为1，但是a的长度和容量不变
	c := append(a, 2)    // a的长度还是0，c=[2], a指向的底层数组的首元素变为2
	fmt.Println(a, b, c) // [] [2] [2]

	a11 := make([]bool, 1)
	a13 := append(a11, true)
	a14 := append(a11, false)
	a11 = append(a11, a13...)
	a11 = append(a11, false)
	fmt.Println(a13, a14, a11)
}
