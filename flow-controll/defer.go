package main

import "fmt"

func main() {
	//ifDeferTest(1)
	deferStackTest()
}

// 会输出
// hello
// world
// defer 放在一个语句之前 可以吧这条语句放在外面
func helloWorld() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// defer 会把这句话放在任何大括号外1层，即便这句话所在域中的变量在外面并不在变量的作用域
func ifDeferTest(x int) {
	if y := 1; x < 10 {
		defer fmt.Println("defer语句 ", y)
		y += 1
		fmt.Println("if代码块中语句 ", y)
	}
}

func deferStackTest() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//返回：
//counting
//done
//9
//8
//7
//6
//5
//4
//3
//2
//1
//0

//说明defer是个栈控制的 先压入的后执行
