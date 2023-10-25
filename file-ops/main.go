package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	fileName := os.Args[1]
	file, er := os.Open(fileName)

	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
