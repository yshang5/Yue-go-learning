package main

import (
	"fmt"
	"math"
)

// compute函数接受一个函数作为参数
// 参数fn的类型是：func(float64, float64) float64
// 即：接受两个float64参数，返回一个float64的函数
// 函数内部调用fn(3, 4)，相当于执行传入的函数
// 等价于：
//
//	public static double compute(BiFunction<Double, Double, Double> fn) {
//		return fn.apply(3.0, 4.0);
//	}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
