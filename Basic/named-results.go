package main

import "fmt"

func main() {
	fmt.Println(split(100))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//
//返回值的命名应当能反应其含义，它可以作为文档使用。
//
//没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值。
//
//裸返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
//
//(别这么用了吧 感觉不好用)
