package main

import "fmt"

func main() {
	//这样会报错，只有接口才能进行类型断言
	//var i = "hello"
	//s, ok := i.(string)
	//fmt.Println(s, ok)

	//下面一句话：空接口可以传值，可以当成java中的object
	var i interface{} = "hello"
	s, ok := i.(string)
	//如果类型命中 s会被赋值为hello
	fmt.Println(s, ok) //hello true

	//如果类型未被命中 integer会被赋值为T的0值 也就是int的0值 0
	integer, ok := i.(int)
	fmt.Println(integer, ok) //0 false
}
