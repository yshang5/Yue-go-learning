package main

import (
	"fmt"
	"time"
)

func fibonacci(limit int, c chan int) {
	x, y := 0, 1
	for i := 0; i < limit; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	//信道可以被遍历
	for i := range c {
		fmt.Println(i)
	}

	time.Sleep(3 * time.Second)
	//当信道被close掉之后, k就会为false 不然就是true，value为缓存的队头
	v, k := <-c
	fmt.Println(v, k)
}
