package main

import (
	"encoding/json"
	"fmt"
)

// Test dsd
type Test struct {
	A *string
	B *string
}

func main() {
	s := "Something"
	t := Test{A: &s, B: nil}
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	copy()
}
