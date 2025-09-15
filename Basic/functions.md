

## Go 语言函数详解

### 1. 基本语法结构

```go
func 函数名(参数列表) 返回值类型 {
    // 函数体
    return 值
}
```

您的 `add` 函数：
```go
func add(x int, y int) int {
    return x + y
}
```

### 2. 与 Java 函数对比

| 方面 | Java | Go |
|------|------|-----|
| **声明** | `public int add(int x, int y)` | `func add(x int, y int) int` |
| **关键字** | `public` + 返回类型 | `func` |
| **参数类型** | 类型在变量名前面 | 类型在变量名后面 |
| **返回** | `return value;` | `return value` |

### 3. Go 函数的多种写法

#### 相同类型参数可以简化
```go
// 您的写法
func add(x int, y int) int {
    return x + y
}

// 简化写法（相同类型）
func add(x, y int) int {
    return x + y
}
```

#### 多返回值（Go 特色）
```go
// 返回多个值
func divide(a, b int) (int, int) {
    return a / b, a % b  // 商和余数
}

// 命名返回值
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return  // 自动返回命名的变量
}
```

#### 无返回值
```go
func printMessage(msg string) {
    fmt.Println(msg)
}
```

#### 可变参数
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 调用：sum(1, 2, 3, 4, 5)
```

### 4. 函数作为值（类似 Java 的 Lambda）

```go
// 函数作为变量
var operation func(int, int) int = add

// 函数作为参数
func calculate(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// 调用
result := calculate(10, 20, add)
```

### 5. 方法（类似 Java 的类方法）

```go
type Calculator struct {
    name string
}

// 方法
func (c Calculator) add(x, y int) int {
    return x + y
}

// 指针接收者方法
func (c *Calculator) setName(name string) {
    c.name = name
}
```

### 6. 匿名函数（类似 Java 的 Lambda）

```go
// 立即执行的匿名函数
result := func(x, y int) int {
    return x + y
}(10, 20)

// 赋值给变量
addFunc := func(x, y int) int {
    return x + y
}
```

### 7. 闭包（Closure）

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 使用
c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
fmt.Println(c()) // 3
```

### 8. 错误处理（Go 特色）

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

// 调用
result, err := divide(10, 0)
if err != nil {
    fmt.Println("错误:", err)
} else {
    fmt.Println("结果:", result)
}
```

## 总结

Go 语言的函数特点：
1. **简洁** - 比 Java 更简洁的语法
2. **多返回值** - 天然支持，无需包装类
3. **一等公民** - 函数可以作为值传递
4. **方法接收者** - 类似 Java 的实例方法
5. **错误处理** - 通过多返回值优雅处理错误

您的 `add` 函数是一个很好的开始，展示了 Go 函数的基本语法！