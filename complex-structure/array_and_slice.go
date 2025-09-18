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

	//1. 数组的切片：从数组创建切片，类似Python的切片操作
	//切片是引用类型，指向原数组的一部分，修改切片会影响原数组
	var spliced []int = primes[1:4]
	fmt.Println(spliced)

	//切片不创建新数组，而是原数组的引用，修改切片数据会改变原数组
	spliced[0] = 6
	fmt.Println(spliced)
	fmt.Println(primes)

	//2. 切片语法：s[开始:结束]，不写开始默认为0，不写结束默认为长度
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[1:4])

	fmt.Println(s[:2])

	fmt.Println(s[1:])

	//3. 切片字面量：[]类型{值}，没有长度限制
	array := []bool{true, true, false}
	fmt.Println(array)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	//4. 结构体切片：可以声明匿名结构体切片
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

	//5. 切片的长度和容量：
	//长度(len)：切片包含的元素个数
	//容量(cap)：从切片第一个元素到底层数组末尾的元素个数
	//可通过len(s)和cap(s)获取

	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 扩展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)

	//输出结果：
	//len=6 cap=6 [2 3 5 7 11 13]  // 原始切片
	//len=0 cap=6 []                // 截取到长度0
	//len=4 cap=6 [2 3 5 7]         // 扩展到长度4
	//len=2 cap=4 [5 7]             // 舍弃前2个元素

	//6. 零值切片：未初始化的切片为零值nil
	var nilArray []int
	fmt.Println(nilArray, len(nilArray), cap(nilArray))
	if nilArray == nil {
		fmt.Println("nil!")
	}

	//7. 使用make创建切片：
	//make([]类型, 长度, 容量) - 容量可选，不填则等于长度
	aSlice := make([]int, 5) // 长度=5，容量=5，元素初始化为0
	printSlice(aSlice)       // 输出：len=5 cap=5 [0 0 0 0 0]

	bSlice := make([]int, 1, 5) // 长度=1，容量=5
	printSlice(bSlice)          // 输出：len=1 cap=5 [0]

	//8. 切片重新切片：可以扩展长度到容量范围内
	cSlice := bSlice[:2] // 重新切片：长度变为2，容量仍为5
	printSlice(cSlice)   // 输出：len=2 cap=5 [0 0]

	//注意：切片操作不能超出当前切片的长度范围
	//dSlice := cSlice[2:5]         // 这会panic！因为cSlice长度只有2
	dSlice := cSlice[1:5] // 正确：从索引1到容量边界
	printSlice(dSlice)    // 输出：len=4 cap=4 [0 0 0 0]

	//9. 切片的追加

	// 可在空切片上追加
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)

	// for range 遍历时 可以在语句中声明两个变量，第一个是下标 第二个是value
	// 如果只声明一个，那就只有值
	iterable := []int{1, 3, 5, 7, 11, 13}
	fmt.Println("下标和值都声明")
	for index, eachValue := range iterable {
		fmt.Printf("下标：%d，值 %d\n", index, eachValue)
	}

	fmt.Println("只声明值")
	for eachValue := range iterable {
		fmt.Printf("下标：%s，值 %d\n", "无", eachValue)
	}

	fmt.Println("下划线代替不需要的声明(下标)")
	for _, eachValue := range iterable {
		fmt.Printf("下标：%s，值 %d\n", "无", eachValue)
	}

	fmt.Println("下划线代替不需要的声明(值)")
	for index, _ := range iterable {
		fmt.Printf("下标：%d，值 %s\n", index, "无")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
