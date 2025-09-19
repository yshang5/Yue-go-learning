package main

import "fmt"

func fibonacci1(dataChannel chan int, quitChannel chan bool) {
	x, y := 0, 1
	for {
		//以下是select的格式 代表着不同的事件
		//select {
		//case <-ch1:           // 事件1：ch1有数据可读
		//case ch2 <- value:    // 事件2：ch2可以写入数据
		//case <-time.After(1): // 事件3：超时事件
		//default:              // 事件4：所有事件都不满足
		//}
		select {

		// 当dataChannel ready去写入数据的时候 往信道里面写x
		case dataChannel <- x:
			x, y = y, x+y
		// 当quitChannel信道里面有数据进来了
		case message := <-quitChannel:
			fmt.Println(message)
			return
		}
	}
}

func main() {
	dataChannel := make(chan int)
	quitChannel := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-dataChannel)
		}
		quitChannel <- true
	}()
	fibonacci1(dataChannel, quitChannel)

}
