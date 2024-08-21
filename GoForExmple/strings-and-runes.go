package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//在 Go 语言中，字符串是一个只读的字节切片。Go 语言和标准库对字符串进行了特殊处理——将其视为以 UTF-8 编码的文本容器。
	//在其他语言中，字符串通常由“字符”组成，而在 Go 语言中，“字符”的概念被称为 rune——它是一个表示 Unicode 代码点的整数。
	//Go 语言官方博客中的一篇文章对这一主题进行了很好的介绍。

	//泰语中的 hello，go语言中是utf-8编码的文本
	const s = "สวัสดี"

	//由于字符串等同于 `[]byte`（字节切片），因此这将产生存储在其中的原始字节的长度。
	fmt.Println("Len:", len(s))

	//对字符串进行索引会生成每个索引位置上的原始字节值。这个循环会生成构成字符串 s 中代码点的所有字节的十六进制值。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	//要计算字符串中有多少个 rune，我们可以使用 utf8 包。
	//需要注意的是，RuneCountInString 的运行时间取决于字符串的大小，因为它需要依次解码每个 UTF-8 rune。
	//一些泰语字符由多个字节组成的 UTF-8 代码点表示，因此这个计数的结果可能会出乎意料。

	//rune 是一种数据类型，用于表示一个 Unicode 代码点。rune 本质上是 int32 的别名，可以用来存储任何有效的 Unicode 字符。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	t := "hello"

	fmt.Println("Rune count t:", utf8.RuneCountInString(t))

	t = "世界"

	fmt.Println("Rune count t:", utf8.RuneCountInString(t))

	//idx 是 for 循环中 range 子句的第一个返回值，表示当前 rune 在字符串 s 中的 起始字节索引。
	//idx 不是逐级递增的， 对应每个rune在utf-8字符串中开始的字节位置，比如第一个ส 这个字符占用了三个字节 所以他后面的 Unicode 代码点是从3开始
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	for idx, runeValue := range t {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	//utf8.DecodeRuneInString我们可以通过明确使用该函数来实现相同的迭代 。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		fmt.Printf("%#U ends at %d\n", runeValue, w)
		fmt.Println(s[i:])
		//width rune 字节长度
		fmt.Println(width)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
