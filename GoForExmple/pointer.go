package main

import (
	"fmt"
	"reflect"
)

// 指针的几种赋值方式
func main() {
	i := 10

	var intPtr *int = &i
	fmt.Println("pointer value :", intPtr, "point to", *intPtr)
	fmt.Println("type of pointer:", reflect.TypeOf(intPtr))

	intPtr2 := &i
	fmt.Println("pointer2 value :", intPtr2, "point to", *intPtr2)
	fmt.Println("type of pointer:", reflect.TypeOf(intPtr2))

	var intPtr3 = &i
	fmt.Println("pointer3 value :", intPtr3, "point to", *intPtr3)
	fmt.Println("type of pointer:", reflect.TypeOf(intPtr3))

	var intPtr4 *int
	intPtr4 = &i
	fmt.Println("pointer4 value :", intPtr4, "point to", *intPtr4)
	fmt.Println("type of pointer:", reflect.TypeOf(intPtr4))

	//不赋值的时候，默认是nil
	var intPtr5 *int
	fmt.Println("pointer5 == nil", intPtr5 == nil)

	//指向数组的指针
	array := [3]int{1, 2, 3}
	var arrayPtr *[3]int = &array
	fmt.Println("array pointer value :", arrayPtr, "point to", *arrayPtr)
	for i := 0; i < len(array); i++ {
		fmt.Printf("arrayPtr[%d]=%d\n", i, (*arrayPtr)[i])
	}

	const SIZE = 5
	//指针数组
	var ptrArray [SIZE]*int
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < SIZE; i++ {
		ptrArray[i] = &a[i]
	}

	for i := 0; i < SIZE; i++ {
		fmt.Printf("%d ", *ptrArray[i])
		fmt.Print(*ptrArray[i], " ")
		fmt.Printf("%d ", ptrArray[i]) // %d 输出十进制形式的地址
		fmt.Printf("%p ", ptrArray[i]) // %p 输出十六进制形式的地址
		fmt.Print(ptrArray[i], "\n")   //%
	}
	fmt.Println()

	a1, b := 1, 2
	swap(&a1, &b)
	fmt.Println("a1=", a1, " b=", b) // a= 2  b= 1

	swap2(&a1, &b)
	fmt.Println("a1=", a1, " b=", b) // a= 2  b= 1
}

// 这个可以交换外部传入的2个实参的值
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// 这个无法交换外部传入的2个实参的值
func swap2(a *int, b *int) {
	//*a, *b = *b, *a
	a, b = b, a
}
