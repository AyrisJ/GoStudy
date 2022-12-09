package main

import "fmt"

func main() {
	// 数组测试
	var a [3]int
	var b [3]int = [3]int{1, 2, 3}

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d,%d\n", i, v)
	}

	for _, v := range b {
		fmt.Printf("%d\n", v)
	}

	// 数组切片
	fmt.Println(b, b[1:2])


}
