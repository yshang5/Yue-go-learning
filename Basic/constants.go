package main

import "fmt"

const Pi = 3.14

func main() {

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// Hello 世界
	// Happy 3.14 Day
	// Go rules? true

	fmt.Printf("Pi is a type of %T", Pi)
	fmt.Println()
	fmt.Printf("Truth is a type of %T", Truth)
	fmt.Println()
	fmt.Printf("World is a type of %T", World)
	fmt.Println()
	// 	常量的声明与变量类似，只不过使用 const 关键字。这点和js很像
	// 常量可以是字符、字符串、布尔值或数值。
	// 常量不能用 := 语法声明。
	// go 的常量也可以做类型推断
}
