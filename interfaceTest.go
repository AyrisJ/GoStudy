package main

import (
	"fmt"
	"sort"
)

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("write data:", data)
	return nil
}

type MyStringList []string

func main() {
	f := new(file)
	var dw DataWriter

	dw = f
	dw.WriteData("hello")

	// 类型判断
	var x interface{}
	x = 10
	vl, ok := x.(int)
	fmt.Println(vl, ",", ok)
	var s1 interface{} = "hello"
	vl2, ok2 := s1.(string)
	fmt.Println(vl2, ",", ok2)

	var a1 int
	a1 = 10
	getType(a1)

	//排序
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	//sort.Sort(names)

	sort.Strings(names)
	for _, v := range names {
		fmt.Println(v)
	}
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
