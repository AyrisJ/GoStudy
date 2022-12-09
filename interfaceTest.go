package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {

}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("write data:", data)
	return nil
}

func main() {
	f:=new(file)
	var dw DataWriter

	dw=f
	dw.WriteData("hello")
}
