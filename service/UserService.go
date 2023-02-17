package service

import (
	"GoStudy/model"
	"fmt"
)

func PrintUserName() string{
	fmt.Printf("print111 %s\n", model.Name)
	return model.Name
}