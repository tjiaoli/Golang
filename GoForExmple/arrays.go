package main

import "fmt"

func main() {

	var a [5]int

	fmt.Println("emp:", a)
	var a1 []int
	var a2 []string
	a2 = append(a2, "1")
	a2 = append(a2, "1", "33", "44")
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("dcl:", b)

	//您还可以让编译器计算元素的数量...
	// [...]先判断了后面的长度是多少，然后赋值给b ，长度不对会报错
	b = [...]int{1, 2, 3, 4, 5, 8, 9}

	fmt.Println("dcl:", b)

	// 此处  6: 8 的意思是 给下标是6的元素赋值8，
	// 但是 如果这样[...]int{8, 3: 8} 编译器会判断最后一个元素是下标3，长度4 与 b不符会报错
	b = [...]int{8, 6: 8}
	fmt.Println("idx:", b)

	e := [...]int{8, 3: 8}
	fmt.Println("dcl:", e)

	//多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
