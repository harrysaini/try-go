package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("abcassde5")

	buf := make([]byte, 5)
	if _, err := io.ReadFull(r, buf); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)

	if _, err := io.ReadFull(r, buf); err != nil {
		fmt.Println("asds", err)
	}
}
