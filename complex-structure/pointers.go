package main

import "fmt"

func main() {
	var p *int
	//* 操作符表示指针指向的底层值。
	//fmt.Println("初始化的时候 p:", p, "指针指向", *p) 这会回有异常“invalid memory address or nil pointer dereference”
	//类似Null pointer
	fmt.Println("初始化的时候 p:", p) //初始化的时候 p: <nil>
	i := 10
	//& 操作符会返回一个指向其操作数的指针
	// in our case 就是 i的指针
	p = &i

	fmt.Println("指向了i   ", "p:", p, "指针指向", *p) //指向了i    p: 0x1400010e020 指针指向 10
	i = 100
	fmt.Println("i换成了100  ", "p:", p, "指针指向", *p)      //i换成了100   p: 0x1400010e020 指针指向 100
	*p = 21                                            // 通过指针 p 设置 i
	fmt.Println("通过指针 p 设置 i后  ", "p:", p, "指针指向", *p) //通过指针 p 设置 i后   p: 0x1400010e020 指针指向 21
	fmt.Println("此时 i: ", i)                           //此时 i:  21

	//总结：修改指针的值，被指向的变量会变， 修改变量，指针也会改变
}
