package main

import (
	"container/list"
	"fmt"
	"sync"
)

var g1, g2 = 100, "global"

func main() {

	var isOk = true
	fmt.Println(isOk)

	var n1, n2 int = 1, 2
	fmt.Println(n1, n2)

	n3 := 3
	fmt.Println(n3)

	fmt.Println(g1, g2)
	var g3 = g2
	fmt.Println(g3)
	fmt.Println(&g1, &g2, &g3, &n1)

	const length, width = 10, 20
	area := length * width
	fmt.Println(area)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		i2 = 1 << iota
		j2 = 3 << iota
		k2
		l2
	)
	fmt.Println("i2=", i2)
	fmt.Println("j2=", j2)
	fmt.Println("k2=", k2)
	fmt.Println("l2=", l2)

	// 追加数组元素
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("len=%d, cap=%d, pointer=%p\n", len(nums), cap(nums), nums)
	}

	// 数组复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2)
	fmt.Println(slice1)
	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{5, 4, 3}
	copy(slice2, slice1)
	fmt.Println(slice2)

	// map集合
	var mapLit map[string]int
	mapLit = map[string]int{"jack": 18, "herry": 20}
	fmt.Println(mapLit["jack"], mapLit["haha"])

	// map遍历
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	var sceneList []string
	for k, v := range scene {
		fmt.Println(k, v)
		sceneList = append(sceneList, k)
	}
	fmt.Println(sceneList)

	delete(scene, "route")
	fmt.Println(scene)

	scene = make(map[string]int)
	fmt.Println(scene)

	// 并发安全的map集合
	var sceneSync sync.Map
	sceneSync.Store("green", 97)
	sceneSync.Store("london", 100)
	sceneSync.Store("egypt", 200)

	fmt.Println(sceneSync.Load("london"))
	sceneSync.Delete("london")
	sceneSync.Range(func(key, value any) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

	list1 := list.New()
	list1.PushBack("canon")
	list1.PushFront(67)
	element := list1.PushBack("second")
	list1.InsertAfter("third", element)
	list1.InsertBefore("first", element)
	fmt.Println(list1)
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("---after remote element")
	list1.Remove(element)
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
