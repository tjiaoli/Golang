package main

import (
	"fmt"
	"maps"
)

func main() {

	//创建映射 make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//初始化 int
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//不存在也能返回值，返回类型的零值
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//返回键值对的数量
	fmt.Println("len:", len(m))

	_, prs1 := m["k2"]
	fmt.Println("prs1:", prs1)

	//删除某个键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	//清空所有
	clear(m)
	fmt.Println("map:", m)

	//从映射（map）中获取值时的可选第二个返回值表示该键是否存在于映射中。这可以用来区分缺失的键和具有零值（如 0 或 ""）的键。
	//在这里，我们不需要实际的值，因此使用空白标识符 _ 忽略了它。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//直接初始化  map[key-type]val-type{key:val,ke:val...}
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	//
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	// == 运算符只允许对数组、指针、接口、通道、函数、和结构体中的某些字段进行比较。
	//if n == n2 {
	//	fmt.Println("n == n2")
	//}
}
