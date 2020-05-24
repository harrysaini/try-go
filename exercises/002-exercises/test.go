package main

import "fmt"

func main() {
	fmt.Print("------- 1 -----\n\n")
	x := 101

	fmt.Printf("%d %b %#x", x, x, x)

	fmt.Println("\n------- *** -----")
	fmt.Print("------- 2 -----\n\n")

	fmt.Println(42 == 42)
	fmt.Println(2 <= 42)
	fmt.Println(42 >= 423)
	fmt.Println(42 != 2)
	fmt.Println(42 < 42)
	fmt.Println(2 > 42)

	fmt.Println("\n------- *** -----")
	fmt.Print("------- 3 -----\n\n")

	const a int = 11
	const b = "harish"
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	// fmt.Println("%v")

	fmt.Println("\n------- *** -----")
	fmt.Print("------- 4 -----\n\n")

	val := 93
	fmt.Println(val)
	fmt.Println(val << 3)
	fmt.Println(val >> 1)

	shifted := val << 1
	lshifted := val >> 1
	fmt.Printf("%d\t%b\t\t%#x\n", val, val, val)
	fmt.Printf("%d\t%b\t%#x\n", shifted, shifted, shifted)
	fmt.Printf("%d\t%b\t\t%#x\n", lshifted, lshifted, lshifted)

	fmt.Println("\n------- *** -----")
	fmt.Print("------- 5 -----\n\n")

	rawString := `dsad
	dsadsad
	dasdsad
	dsfsfsf
	fdsfd
	s`

	fmt.Println(rawString)

	fmt.Println("\n------- *** -----")
	fmt.Print("------- 6 -----\n\n")

	const (
		one = iota + 2015
		two
		three
		four
	)

	const a1, b1 = 89, 89

	fmt.Println(one, two, three, four)
	fmt.Println(a1, b1)

}
