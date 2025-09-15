package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
}

//每个 Go 程序都由包构成。
//
//程序从 main 包开始运行。
//
//本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
//按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。

//package main = Java 中声明"这个类包含 main 方法"
//func main() = Java 的 public static void main(String[] args)
//只有声明了 package main 的包才能编译成可执行文件
//其他包名（如 package utils）相当于 Java 中的普通类，只能被导入使用
//所以您的理解完全正确！package main 确实相当于告诉 Go 编译器"这个包包含程序入口点"。
