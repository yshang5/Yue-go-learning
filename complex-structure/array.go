package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//数组的切片 很像python的list的用法
	var spliced []int = primes[1:4]
	fmt.Println(spliced)

	//实际上切片并不返回一个新的数组 而是原来数组的引用，当修改切片的数据，原数组也会改变
	spliced[0] = 6
	fmt.Println(spliced)
	fmt.Println(primes)

	//切片的用法和python一样 前面的不写代表0 后面的代表length
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[1:4])

	fmt.Println(s[:2])

	fmt.Println(s[1:])

	//这是一个没有长度的字面量
	array := []bool{true, true, false}
	fmt.Println(array)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	//structure数组 可以这样声明 （不需要type关键字）
	structures := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(structures)
}
