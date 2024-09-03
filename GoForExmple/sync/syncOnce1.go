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
	s2.SetData("second call")
	fmt.Println(s2.data) // 输出: first call, 数据未被覆盖
}
