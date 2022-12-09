package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fileName := "simple.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}

	defer file.Close()

	var chunk []byte
	buf := make([]byte, 1024)
	r := bufio.NewReader(file)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)

	}

	fmt.Println("read context:")
	fmt.Println(string(chunk))
}
