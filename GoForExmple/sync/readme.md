`sync.Once` 的 `Do` 方法用于确保某个函数只执行一次，但 `Do` 方法本身不能直接传递参数给函数。因此，如果你需要在被 `sync.Once` 控制的函数中使用参数，有以下几种常见的解决方案：

### 1. **使用闭包（匿名函数）捕获参数**
   - 可以通过闭包将参数捕获到函数的上下文中，然后再调用 `sync.Once` 的 `Do` 方法。

```go
package main

import (
	"fmt"
	"sync"
)

var instance *Singleton
var once sync.Once

type Singleton struct {
	data string
}

func GetInstance(data string) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: data, // 使用传入的参数
		}
	})
	return instance
}

func main() {
	s1 := GetInstance("first call")
	fmt.Println(s1.data) // 输出: first call

	s2 := GetInstance("second call")
	fmt.Println(s2.data) // 输出: first call, 单例模式确保初始化只执行一次
}
```

### 2. **预先存储参数**
   - 如果需要在初始化时使用的参数是已知的，或者可以提前获取，可以先将参数存储在变量中，然后在 `sync.Once` 的 `Do` 方法中使用。

```go
package main

import (
	"fmt"
	"sync"
)

var instance *Singleton
var once sync.Once
var initData string // 预先存储参数

type Singleton struct {
	data string
}

func initSingleton() {
	instance = &Singleton{
		data: initData,
	}
}

func GetInstance(data string) *Singleton {
	initData = data
	once.Do(initSingleton)
	return instance
}

func main() {
	s1 := GetInstance("first call")
	fmt.Println(s1.data) // 输出: first call

	s2 := GetInstance("second call")
	fmt.Println(s2.data) // 输出: first call, 单例模式确保初始化只执行一次
}
```

### 3. **懒惰初始化后再设置参数**
   - 在这种模式下，`sync.Once` 只控制实例的初始化，而参数设置可以通过另一个方法进行。如果参数需要频繁改变，这种模式更合适。

```go
package main

import (
	"fmt"
	"sync"
)

var instance *Singleton
var once sync.Once

type Singleton struct {
	data string
}

func initSingleton() {
	instance = &Singleton{}
}

func GetInstance() *Singleton {
	once.Do(initSingleton)
	return instance
}

func (s *Singleton) SetData(data string) {
	s.data = data
}

func main() {
	s1 := GetInstance()
	s1.SetData("first call")
	fmt.Println(s1.data) // 输出: first call

	s2 := GetInstance()
	fmt.Println(s2.data) // 输出: first call, 数据未被覆盖
}
```

### 总结

- **闭包捕获参数**：这是最常用的方法，可以直接在 `Do` 方法中通过闭包使用外部参数。
- **预先存储参数**：适用于参数在调用前就已经知道的情况。
- **懒惰初始化后再设置参数**：适用于实例初始化和参数设置分开的情况，灵活性更高。
