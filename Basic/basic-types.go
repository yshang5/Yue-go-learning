// go 语言的一些基本类型
package main

import "fmt"

//这个斜线 / 是 Go 语言中包路径分隔符
//math/cmplx 是 Go 语言cmplx 是 Go 语言中专门用于复数运算的包，cmplx 是 "complex" 的缩写。
import "math/cmplx"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("类型：%T 值：%v\n", ToBe, ToBe)
	fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)
	fmt.Printf("类型：%T 值：%v\n", z, z)
}

// bool - 布尔类型，只有 true 和 false 两个值

// string - 字符串类型，不可变字符串

// int  int8  int16  int32  int64 - 有符号整数类型
// int: 平台相关，32位或64位
// int8: 8位有符号整数，范围 -128 到 127
// int16: 16位有符号整数，范围 -32768 到 32767
// int32: 32位有符号整数，范围 -2^31 到 2^31-1
// int64: 64位有符号整数，范围 -2^63 到 2^63-1

// uint uint8 uint16 uint32 uint64 uintptr - 无符号整数类型
// uint: 平台相关，32位或64位
// uint8: 8位无符号整数，范围 0 到 255
// uint16: 16位无符号整数，范围 0 到 65535
// uint32: 32位无符号整数，范围 0 到 2^32-1
// uint64: 64位无符号整数，范围 0 到 2^64-1
// uintptr: 指针类型，用于存储指针的整数

// byte - uint8 的别名，用于表示字节数据

// rune - int32 的别名，表示 Unicode 码点，支持完整 Unicode

// float32 float64 - 浮点数类型
// float32: 32位浮点数，类似 Java 的 float
// float64: 64位浮点数，类似 Java 的 double

// complex64 complex128 - 复数类型
// complex64: 64位复数，实部和虚部各32位
// complex128: 128位复数，实部和虚部各64位

// 本例展示了几种类型的变量。 和导入语句一样，变量声明也可以「分组」成一个代码块。

// int、uint 和 uintptr 类型在 32-位系统上通常为 32-位宽，在 64-位系统上则为 64-位宽。当你需要一个整数值时应使用 int 类型， 除非你有特殊的理由使用固定大小或无符号的整数类型。
