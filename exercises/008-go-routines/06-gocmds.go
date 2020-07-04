package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS\t", runtime.GOOS)
	fmt.Println("GOARCH\t", runtime.GOARCH)
}
