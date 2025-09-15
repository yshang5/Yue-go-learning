package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	//1 2 true false no!
}

//变量的初始化
//变量声明可以包含初始值，每个变量对应一个。
//
//如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。
