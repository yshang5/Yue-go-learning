package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	// 方式1：按顺序初始化
	v1 := Vertex{1, 2}

	// 方式2：指定字段名
	v2 := Vertex{X: 1, Y: 2}

	// 方式3：部分初始化
	v3 := Vertex{X: 1} // Y默认为0

	var v4 = Vertex{}
	// 方式4：使用new
	v5 := &v4 // 返回指针

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(*v5)
	fmt.Println(v1 == v2)
	fmt.Println(&v1 == &v2)
	fmt.Println(&v4 == v5)

}
