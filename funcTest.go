package main

import (
	"errors"
	"flag"
	"fmt"
	"time"
)

func hypot(x, y float64) float64 {
	return x * y
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func typedTwoValues() (int, int) {
	return 1, 2
}

func fire() {
	fmt.Println("fire")
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

var skillParam = flag.String("skill", "fire2", "skill to perform")

type Invoker interface {
	Call(interface{})
}
type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

// 递归函数调用
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {
	// 函数调用
	fmt.Println(hypot(3, 4))

	// 函数定义
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

	// 函数返回值
	a, b := typedTwoValues()
	fmt.Println(a, b)

	// 函数变量引用赋值调用
	var f func()
	f = fire
	f()

	// 匿名函数
	f2 := func(data int) {
		fmt.Println("hello", data)
	}
	f2(100)

	// lanmda表达式
	visit([]int{1, 2, 3, 4, 5}, func(i int) {
		fmt.Print(i)
	})
	fmt.Println()

	// 传入参数自动选择函数处理
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

	// 函数类型实现接口
	var invoker Invoker
	s := new(Struct)

	invoker = s
	invoker.Call("Hello")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function:", v)
	})
	invoker.Call("Hello")

	//闭包测试
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)

	// 可变参数类型
	myfunc := func(args ...int) {
		for _, arg := range args {
			fmt.Print(arg)
		}
	}
	myfunc(2, 3, 4)
	fmt.Println()
	myfunc(1, 5, 6, 8, 9)

	// interface传递任意数据
	myPrint := func(args ...interface{}) {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
			case string:
				fmt.Println(arg, "is a string value.")
			case int64:
				fmt.Println(arg, "is an int64 value.")
			default:
				fmt.Println(arg, "is an unknown type.")
			}
		}
	}

	myPrint(1, 234, "hello", 1.234)

	// 递归函数测试
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	// 异常处理
	// 定义除数为0的错误
	var errDivisionByZero = errors.New("division by zero")
	divideFun := func(upNum, downNum int) (int, error) {
		if downNum == 0 {
			return 0, errDivisionByZero
		}
		return upNum / downNum, nil
	}
	fmt.Println(divideFun(1, 0))

	// 程序崩溃模拟
	//panic("crash")

	// 计算函数执行耗时
	timeCostFunc := func() {
		start := time.Now()
		sum := 0
		for i := 0; i < 100000000; i++ {
			sum++
		}
		elapsed := time.Since(start)
		fmt.Println("执行耗时:", elapsed)
	}
	timeCostFunc()
}
