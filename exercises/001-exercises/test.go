package main

import "fmt"

func main() {
	// exerciseOne();
	// exerciseTwo();
	exerciseThree()
	exerciseFour()
}

func exerciseOne() {
	fmt.Printf("\n\n---- 1 ----- \n")

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
}

// var x int
// var y string
// var z bool

func exerciseTwo() {
	fmt.Printf("\n\n---- 2 ----- \n")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

var x int = 42
var y string = "James Bond"
var z bool

func exerciseThree() {
	fmt.Printf("\n\n---- 3 ----- \n")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)

}

type myInt int

var in2 myInt
var in3 int

func exerciseFour() {
	fmt.Printf("\n\n---- 4 ----- \n")

	fmt.Println(in2)
	fmt.Printf("%T\n", in2)
	fmt.Println(in3)
	fmt.Printf("%T\n", in3)

	in2 = 12

	str1 := "saasdsc"

	fmt.Println(str1)

	fmt.Println(str1)

	fmt.Println(in2)

	in3 = int(in2)

	fmt.Println(in3)

	fmt.Printf("\n\n---- STrings ----- \n")

	fmt.Println(len("Hello, World"))
	fmt.Printf("%c\n", "Hello, World"[1])
	fmt.Println("Hello, " + "World")

	fmt.Println(32132 * 42452)

}
