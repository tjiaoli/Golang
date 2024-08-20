package main

import (
	"fmt"
	"slices"
)

// 切片
func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//要创建非零的空切片，请使用内置的make。这里我们创建一个string长度为3的切片s，默认情况下，新切片的容量等于其长度，长度需要显示参数给make
	s = make([]string, 3)
	//len长度  cap容量
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	//我们可以像数组一样设置和获取。
	s[0] = "a"
	//s[1] = "b" 虽然注释了没有复制，但是会占用一个位置，
	s[2] = "c"
	//s[3] = "d"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//append返回一个或多个新增的切片， 此操作是在后面追加元素
	// 创建一个长度为2，容量为3的切片
	s1 := make([]int, 2, 3)
	fmt.Println(s1, len(s1), cap(s1)) // 输出: [0 0] 2 3

	// 追加一个元素，切片的长度增加，容量保持不变
	s1 = append(s1, 1)
	fmt.Println(s1, len(s1), cap(s1)) // 输出: [0 0 1] 3 3

	// 再次追加元素，这次容量不够，append 会创建新的数组
	s1 = append(s1, 2)
	fmt.Println(s1, len(s1), cap(s1)) // 输出: [0 0 1 2] 4 6

	//切片s: [0 0], len: 2, cap: 2, 地址: 0xc000014080
	//切片s: [0 0 3], len: 3, cap: 4, 地址: 0xc0000180c0
	//Go语言有自动垃圾回收机制（Garbage Collector, GC），会自动管理内存。
	//具体机制：
	//引用计数：在Go中，切片是通过指针引用底层数组的。如果没有任何变量引用原来的底层数组，这个数组就会被视为“不可达”的数据。
	//垃圾回收：Go的垃圾回收器会定期扫描内存，发现不再被引用的内存块（比如这个没有引用的数组），然后释放这些内存，使得内存可以被重新分配给其他对象。
	//因此，当你使用append创建了一个新的数组并且原来的数组不再被任何变量引用时，那个原来的数组最终会被垃圾回收，不会造成内存泄漏。
	//这是自动垃圾回收机制的一个重要优势，它让开发者不必手动管理内存释放。

	s = append(s, "d")
	// s原来的容量是3 现在超过了原来的容量会新建一个数组将原来的数据复制到现在的地址
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	//切片可以被复制
	copy(c, s)
	fmt.Println("cpy:", c)
	//切分 s[x:y] 包括x 不包括 y
	l := s[2:5]
	fmt.Println("sl1:", l)
	// y-1
	l = s[:5]
	fmt.Println("sl2:", l)
	// x 到 len - 1
	l = s[2:]
	fmt.Println("sl3:", l)

	//直接声明初始化元组
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	fmt.Println("dcl:", len(t), cap(t))

	t2 := []string{"g", "h", "i"}
	//内容比较
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	//slices 是引用类型 不能用 == 运算符只允许对数组、指针、接口、通道、函数、和结构体中的某些字段进行比较。
	/*if t == t2 {
		fmt.Println("t == t2")
	}*/
	t3 := [...]string{"2", "3"}

	t4 := [...]string{"2", "3"}

	if t3 == t4 {
		fmt.Println("t3 == t4")
	}

	//切片可以组成多维数据结构。与多维数组不同，内部切片的长度可以变化。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//多维切片的 append 操作
	//多维切片本质上是一个切片的切片，因此 append 操作只作用于其中一维。例如，对于一个二维切片，可以向其添加一维切片。
	// 创建一个二维切片
	t5 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("原始切片:", t5)

	// 添加一个新的一维切片
	t5 = append(t5, []int{7, 8, 9})
	fmt.Println("追加后:", t5)

	// 向第一维的第一个元素添加新元素
	t5[0] = append(t5[0], 10)
	fmt.Println("修改后:", t5)
	//每一维度都是一个独立的切片：可以对任意一维的切片进行 append 操作，添加新的元素或子切片。
	//容量变化：和普通切片一样，如果追加的元素使切片的容量不足，Go会自动分配新的底层数组并拷贝原来的数据。
}
