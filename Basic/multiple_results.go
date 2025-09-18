package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

//像python一样可以返回两个值，每个入参和出参都需要规定数据类型
//a := 10 相当于var a int = 10
//var name string = "张三" 相当于 name := "张三"
