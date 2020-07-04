package main

import "fmt"

func main() {
	b := [3]byte{'a', 'b', 'c'}

	fmt.Println(len(b), cap(b))

	b = append(b, []byte("RAM JI KI JAI HO")...)
	fmt.Println(len(b), cap(b))
	fmt.Println(b)
}
