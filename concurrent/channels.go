package main

import (
	"fmt"
	"time"
)

func calculateSum(slice []int, c chan int) {
	sum := 0
	for value := range slice {
		if value < 0 {
			time.Sleep(100 * time.Millisecond)
		}
		sum += value
	}
	//把结果放进通道里，到时候使用 <- c 能把结果倒出来
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	//chan 的初始化 使用make 参数是chan chan的泛型模板，这里是int
	//make(chan int)这种只有一个参数的不带缓冲区，当缓冲区为空时，调用者方会阻塞。
	//make(chan int, 200) 开辟200长度的缓存
	c := make(chan int)
	//chan可以同时被多个线程使用
	//线程1处理前半部
	go calculateSum(s[:len(s)/2], c)
	//线程2处理后半段
	go calculateSum(s[len(s)/2:], c)
	//谁先结束 x就是谁，并且阻塞等待到两个信道的结果全部返回
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
