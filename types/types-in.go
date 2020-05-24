package main

import (
	"fmt"
	"runtime"
)

var a int

type hotdog int
type hot byte

var b hotdog
var c hot

func main() {
	a = 42
	b = 43
	c := 'a'
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	a := int(c)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("------- STRINGs -----")

	name := "Hello, 世界,-- जैसा"

	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)
	fmt.Printf("%x\n", name)

	bs := []rune(name)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U -> %v\n", bs[i], bs[i])
	}

	fmt.Print("\n")

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%b ", bs[i])
	}
	fmt.Println()

	// unicode to string
	ya := 0x091C
	fmt.Printf("%s", string(ya))

	fmt.Println("\n------- *** -----")

	fmt.Println("\n\n------- CONSTANTS -----")

	const a1 int = 10

	fmt.Print(a1, "\n")

	// a1 = 11;

	fmt.Println("\n------- *** -----")

	fmt.Println("\n------- iota -----")

	const (
		i = iota*2 + 12*iota
		j
		k
	)

	fmt.Println(i, j, k)

	const (
		l = iota + 2
		m = 56 + iota
		n
	)

	fmt.Println("---->", l, m, n)

	// iota increments from 0, like enum in typescript kind of
	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)

	fmt.Printf("%v\t\t%b\n", kb, kb)
	fmt.Printf("%v\t\t%b\n", mb, mb)
	fmt.Printf("%v\t\t%b\n", mb, mb)

	// as := "a"
	// int(as)++

	fmt.Print(3 % 1)

}
