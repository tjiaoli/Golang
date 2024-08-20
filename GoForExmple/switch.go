package main

import (
	"fmt"
	"time"
)

// 与java有差别
func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Saturday)
	switch time.Now().Weekday() {
	//您可以使用逗号分隔同一case语句中的多个表达式。我们 default在此示例中也使用了可选的情况。
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch不使用表达式是表达 if/else 逻辑的另一种方式。这里我们还展示了表达式如何 case可以是非常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	case t.Day() > 5:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's after noon")
	}

	//类型switch比较的是类型而不是值。您可以使用它来发现接口值的类型。在此示例中，变量t将具有与其子句相对应的类型

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
