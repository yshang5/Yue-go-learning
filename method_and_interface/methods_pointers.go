package main

import "fmt"

type Counter struct {
	value int
}

// Add 值接收者：不能修改原值
func (c Counter) Add() {
	c.value++ // 只修改副本
	fmt.Printf("值接收者中: %d\n", c.value)
}

// 指针接收者：可以修改原值
// 如果接收是个指针，那么结构体里面的值会被修改，相比于上面的值接收者，只取了个副本并且吧副本里的值改了
func (c *Counter) AddPointer() {
	c.value++ // 修改原值
	fmt.Printf("指针接收者中: %d\n", c.value)
}

func main() {
	counter := Counter{value: 0}

	fmt.Printf("初始值: %d\n", counter.value)

	counter.Add()                          // 值接收者
	fmt.Printf("调用后: %d\n", counter.value) // 还是0！

	counter.AddPointer()                   // 指针接收者
	fmt.Printf("调用后: %d\n", counter.value) // 变成1了！
}
