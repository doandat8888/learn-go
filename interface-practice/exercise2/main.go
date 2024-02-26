package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Os args: ", os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
