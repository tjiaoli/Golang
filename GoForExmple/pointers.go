package main

import "fmt"

//指针

func zeroval(ival int) {
	ival = 0
}

// 当你调用 zeroptr(&i) 时，传递的是 i 的地址，iptr 指向 i 的内存位置。通过 *  = 0，你直接修改了 i 的值。
// 固定写法  val *type  { *val = val }   funcName(&val)
// *int：表示一个指向 int 类型的指针。
// *iptr：对指针 iptr 进行解引用，访问或修改指针指向的变量的值。 对应下面传参 固定写法 &i
func zeroptr(iptr *int) {
	*iptr = 0
}

const SIZE = 5

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	var ptrArray [SIZE]*int
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < SIZE; i++ {
		ptrArray[i] = &a[i]
	}

	for i := 0; i < SIZE; i++ {
		fmt.Printf("%d ", *ptrArray[i])
	}
	fmt.Println()

	//指向指针的指针
	var a1 int = 100
	var ptr1 *int = &a1
	var ptr2 **int = &ptr1
	var ptr3 ***int = &ptr2

	fmt.Println("*ptr1=", *ptr1)
	fmt.Println("**ptr2=", **ptr2)
	fmt.Println("***ptr3=", ***ptr3)

}
