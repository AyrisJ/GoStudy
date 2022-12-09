package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// 基础定义
	type Player struct {
		name        string
		healthPoint int
		magicPoint  int
	}

	p1 := new(Player)
	p1.name = "jack"
	p1.healthPoint = 300

	fmt.Println(p1, &p1)

	p2 := &Player{}
	p2.name = "harry"
	p2.healthPoint = 150
	fmt.Println(p2, &p2)

	// 含指针的结构体
	type Command struct {
		name    string
		var2    *int
		comment string
	}

	var version = 1
	cmd := &Command{}
	cmd.name = "version"
	cmd.var2 = &version
	cmd.comment = "show version"
	fmt.Println(cmd, &cmd)

	// 函数生成结构体
	newCommand := func(name string, varref *int, comment string) *Command {
		return &Command{
			name:    name,
			var2:    varref,
			comment: comment,
		}
	}

	cmd2 := newCommand("v2", &version, "show v2")
	fmt.Println(cmd2, &cmd2)

	// 初始化结构体
	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)

	// 结构体派生
	type Cat struct {
		Name  string
		Color string
	}

	type BlackCat struct {
		Cat
	}
	newCat := func(name string) *Cat {
		return &Cat{
			Name: name,
		}
	}
	newBlackCat := func(color string) *BlackCat {
		cat := &BlackCat{}
		cat.Color = color
		return cat
	}

	fmt.Println(newCat("kimmy"))
	fmt.Println(newBlackCat("black color"))
	fmt.Println(Cat{Name: "tom", Color: "yellow"})

	// 结构体组合
	type Wheel struct {
		Size int
	}

	type Engine struct {
		Power int
		Type  string
	}

	type Car struct {
		Wheel
		Engine
	}

	car1 := Car{
		// 初始化轮子
		Wheel: Wheel{Size: 50},
		// 初始化引擎
		Engine: Engine{Power: 250, Type: "1.4T"},
	}

	fmt.Println(car1)

	// 结构体模拟遍历
	type Node struct {
		data int
		next *Node
	}

	printNodeFunc := func(p *Node) {
		for p != nil {
			fmt.Println(*p)
			p = p.next
		}
	}

	head := &Node{data: 1}
	node1 := &Node{data: 2}
	node2 := &Node{data: 3}
	head.next = node1
	node1.next = node2
	printNodeFunc(head)

	// io读写
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	rf := bufio.NewReader(rd)
	var buf [128]byte
	n, err := rf.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	n2, err2 := w.WriteString("c语言中文网")
	w.Flush()
	fmt.Println(string(wr.Bytes()), n2, err2)

}
