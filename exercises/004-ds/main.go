package main

import "fmt"

func main() {

	fmt.Println("---> Array 1")

	var x [5]int

	x[1] = 10

	fmt.Println(x)

	y := [5]int{1, 2, 3, 4, 5}

	// y[6] = 1

	fmt.Printf("%T\t%v\n", y, y)

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println("---> slice 2")

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\t%v\n", xi, xi)

	for i, v := range xi {
		fmt.Println("elem at", i, v)
	}

	fmt.Println(len(xi), cap(xi))

	xi = append(xi, 12)
	fmt.Println(len(xi), cap(xi))

	fmt.Println("---> slice 3")
	yi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(yi[:])
	fmt.Println(yi[len(yi)-1:])

	more()

}

func more() {
	fmt.Println("---> slice 4")
	var y []int
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y = append(x, 52)
	fmt.Println(y)
	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 55, 53, 54)
	fmt.Println(x)

	z := []int{56, 57, 58, 59, 60}

	x = append(x, z...)
	fmt.Println(x)

	deleteSlice()

}

func deleteSlice() {
	fmt.Println("---> slice 4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:4], x[6:]...)
	fmt.Println(y)

	makeSlice()
}

func makeSlice() {
	fmt.Println("---> make 4")
	x := make([]string, 50, 50)
	fmt.Println(x, len(x), cap(x))

	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	twoDSlice()
}

func twoDSlice() {
	fmt.Println("---> 2d slice 4")
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{
		jb,
		mp,
	}

	fmt.Println(x)

	for _, v := range x {
		for _, st := range v {
			fmt.Println(st)
		}
	}

	maps()
}

func maps() {

	fmt.Println("---> maps 4")

	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(x)

	for key, val := range x {
		fmt.Printf("%v\t%v\n", key, val)
	}

	x[`shaktiman`] = []string{`dr jackel`, `geeta`, `juice`}

	fmt.Println()
	for key, val := range x {
		fmt.Printf("%v\t%v\n", key, val)
	}

	if v, ok := x["shaktiman"]; ok {
		fmt.Println("founf", v)
	}

	delete(x, "shaktiman")
	fmt.Println()
	for key, val := range x {
		fmt.Printf("%v\t%v\n", key, val)
	}

}
