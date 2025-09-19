package main

import "fmt"

// 此声明意味着 s 是满足内置约束 comparable 的任何类型 T 的切片。 x 也是相同类型的值。
// comparable 是一个有用的约束，它能让我们对任意满足该类型的值使用 == 和 != 运算符。在此示例中，我们使用它将值与所有切片元素进行比较，直到找到匹配项。 该 Index 函数适用于任何支持比较的类型。
func Index[T comparable](s []T, e T) int {
	for index, value := range s {
		if value == e {
			return index
		}
	}
	return -1
}

func main() {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
