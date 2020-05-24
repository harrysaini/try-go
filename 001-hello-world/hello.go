package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello world", 43, "ads")
	fmt.Println(n, e)
	variables()
	variable2()
	types()
	fmtTest()
}

// short hand variable
func variables() {
	fmt.Print("\n\n---- SHORT DECLERATION VARIABLES ----- \n")
	x := 10
	y := x * 78
	fmt.Println(x, y)

}

func variable2() {
	fmt.Print("\n\n---- VARIABLES ----- \n")

	// other way to declare
	var x = 43
	fmt.Println(x)

	// unassigned decleration
	// sets default value(ZERO Value) for type defined

	var y int

	fmt.Println(y)

	y = 12
	fmt.Println(y)

}

func types() {
	fmt.Print("\n\n---- Types ----- \n")

	var x = 43
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y string

	fmt.Println(y)

	y = `ds raw  """" string 
	 same t same
	`
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}


func fmtTest() {
	fmt.Printf("\n\n---- FMT ----- \n")

	y := 911

	fmt.Printf("%x\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%#x\n", y)

}

func loops() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}
