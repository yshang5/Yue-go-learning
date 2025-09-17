package main

import "fmt"

func main() {
	for1()
	for2()
}

func for1() {
	sum := 0
	//0到9以内加法
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func for2() {
	sum := 0
	i := 0

	// for语句之前，游标可以初始化
	// 要么完整形式：for 初始化; 条件; 后置语句 { }
	// 要么只能是条件 后置语句放置在块内

	//此时 这个for就按照其他语句中的while，也就是 for 判定 {代码块}
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
