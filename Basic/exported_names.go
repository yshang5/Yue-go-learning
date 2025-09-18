package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

//Pi 是math包中的名字

// 如果是对外的变量或者是函数，首字母必须大写 不然无法被其他包访问
// 1. 导出规则
// 首字母大写 = 导出（public）
// 首字母小写 = 未导出（private）
// 2. 与 Java 对比
// Go	Java	说明
// var PublicVar	public int publicVar	导出的变量
// var privateVar	private int privateVar	未导出的变量
// func PublicFunc()	public void publicFunc()	导出的函数
// func privateFunc()	private void privateFunc()	未导出的函数
// 对外的变量或者是函数，首字母必须大写 不然无法被其他包访问
