package main

import "fmt"

func main() {
	v := 42 // 修改这里看看！

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("v is of type %T\n", v)
	// v is of type int
	// 这里 v是int类型 因为42是数字中最高优先级，尽管v可以是其他类型比如int8
}
