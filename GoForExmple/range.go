package main

import "fmt"

func main() {

	//range 可以用于多种数据结构的元素迭代。让我们来看一下如何在我们已经学习过的一些数据结构中使用 range
	nums := []int{2, 3, 4}
	sum := 0
	//第一个返回值是索引（_ 表示我们不需要这个值，所以使用空白标识符 _ 丢弃它）。
	//第二个返回值是元素值 num，我们在循环中使用它来累加到 sum。
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//如果没有用_,那num只会获取索引值'0,1,2'
	for num := range nums {
		sum += num
	}
	fmt.Println(sum)
	//这里i就被赋予了索引值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	//k v的迭代
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 迭代获取k
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//用于字符串时，会遍历字符串中的 Unicode 码点。第一个返回值是 rune（字符）在字符串中起始字节的索引，
	//第二个返回值是该 rune 本身。有关更多细节，请参阅字符串和 rune。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
