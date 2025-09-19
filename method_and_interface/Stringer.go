package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 隐式实现了Stringer接口，Stringer内容：
//
//	type Stringer interface {
//		String() string
//	}
//
// 接收参数如果是T 并且重写了String方法
// 所有的Stringer的实现在打印的时候都是String()方法返回的内容，很像toString()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
