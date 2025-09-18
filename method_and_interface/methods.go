package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// func关键字和方法名之间是(v Vertex)
// 这个是方法接收者 表示这个方法是属于Vertex类型的
// 所以后面调用的时候是v.Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// 在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。
// 你只能为在同一个包中定义的接收者类型声明方法，而不能为其它别的包中定义的类型 （包括 int 之类的内置类型）声明方法。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
