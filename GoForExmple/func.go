package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

/*
changeRadius和changeRadius2的区别是后者可以改变变量c的成员radius的值，前者不能改变
*/
func (c Circle) changeRadius(radius float64) {
	c.radius = radius
	//c.radius是一个副本生命周期随着函数结束结束
	fmt.Println("radiusbak=", c.radius, "areabak=", c.getArea()) //10, 314
}

func (c *Circle) changeRadius2(radius float64) {
	c.radius = radius
}

func (c Circle) addRadius(x float64) float64 {
	return c.radius + x
}

func main() {
	var c Circle
	c.radius = 10
	fmt.Println("radius=", c.radius, "area=", c.getArea()) //10, 314

	c.changeRadius(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea()) //10, 314

	c.changeRadius2(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea()) //20, 1256

	result := c.addRadius(3.6)
	fmt.Println("radius=", c.radius, "result=", result) // 20, 23.6
}
