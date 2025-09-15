package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("现在你有了 %g 个问题。\n\n", math.Sqrt(7))
}

//此代码用圆括号将导入的包分成一组，这是“分组”形式的导入语句。
//
//当然你也可以编写多个导入语句，例如：
//
//import "fmt"
//import "math"
//不过使用分组导入语句要更好。
