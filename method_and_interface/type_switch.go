package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("I don't know about type that")
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
