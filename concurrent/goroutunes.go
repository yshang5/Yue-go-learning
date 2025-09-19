package main

import "fmt"
import "time"

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//go 就相当于 new Thread(() -> say("world")).start()
	go say("world")
	fmt.Println("Hello, Reader!")
	//主线程直接调用
	say("hello")
}
