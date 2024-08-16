package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	//赋值定义了显示形式
	//const d = 300000000000000000000000000 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	//math.sin 角度的正弦值 单位默认是float64
	//常量表达式执行具有任意精度的算术运算。
	//数字常量没有类型，除非通过显式转换等方式指定类型。
	//通过在需要类型的上下文中使用数字，例如变量赋值或函数调用，可以为其指定类型。例如，此处 math.Sin需要float64。

	fmt.Println(math.Sin(n))

	angle := math.Pi / 4 //45°转化为弧度   math.pi 圆周率Π，
	fmt.Println(angle)
	//角度与弧度的关系是：1度 = π/180 弧度。
	//因此，45 度的弧度可以通过以下计算得到
	//45 度 = 45 * π/180 = π / 4
	sineValue := math.Sin(angle) //45°的正弦值
	fmt.Println("45度的正弦值为:", sineValue)

}
