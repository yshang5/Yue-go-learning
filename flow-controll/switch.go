package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switchWithoutCondition()
}

func naiveSwitch() {
	fmt.Print("Go 运行的系统环境：")
	//和 if 一样可以赋值或者带一句话
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchWithoutCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
