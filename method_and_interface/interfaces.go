package main

import (
	"fmt"
	"math"
)

type abser interface {
	Abs() float64
}

type vertex struct {
	X, Y float64
}

type myFloat float64

// 为myFloat实现Abs方法
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 为*Vertex实现Abs方法
func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a abser

	// 测试myFloat
	f := myFloat(-math.Sqrt2)
	a = f
	fmt.Printf("myFloat: %.2f\n", a.Abs())

	// 测试*Vertex
	v := vertex{3, 4}
	a = &v // 注意：需要指针
	fmt.Printf("Vertex: %.2f\n", a.Abs())
}
