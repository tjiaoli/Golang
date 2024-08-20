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

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
