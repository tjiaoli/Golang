在 Go 语言中，`for` 是唯一的循环结构，用于创建各种形式的循环。它可以用于传统的条件循环、范围循环、以及无限循环。

### 创建 `loop` 的方式：

#### 1. 传统的 `for` 循环（类似于其他语言的 `while` 循环）：
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```
- 这段代码会输出 `0 1 2 3 4`，其中 `i++` 每次循环后自增，直到 `i < 5` 条件为 `false` 时停止。

#### 2. 无限循环：
```go
package main

import "fmt"

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        if i == 5 {
            break // 使用break来结束循环
        }
    }
}
```
- 这里的 `for` 没有条件，表示无限循环，直到通过 `break` 语句主动跳出。

#### 3. 类似 `while` 的循环：
```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```
- 这种形式与其他语言中的 `while` 类似，`for` 后面直接跟条件表达式，循环会持续执行，直到条件为 `false`。

#### 4. `range` 循环（用于遍历集合、数组、切片、字符串等）：
```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```
- `range` 用于遍历集合。上例遍历了一个切片 `nums`，并输出每个元素的索引和值。

### 总结：
- `for` 是 Go 中唯一的循环关键字，可以根据需要创建不同形式的循环：条件循环、无限循环和集合循环。
