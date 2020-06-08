// package main

// import "fmt"

// func main() {
// 	func() {
// 		fmt.Println("I am annoymous")
// 	}()

// 	x := func() {
// 		fmt.Println(" I am variable func")
// 	}

// 	x()

// 	z := foo(12)

// 	fmt.Println(z(1))
// 	fmt.Println(z(2))
// 	fmt.Println(z(3))

// 	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	dbls := listMapper(xi, func(i int, v int) int {
// 		return v * 2
// 	})

// 	fmt.Println(dbls)

// 	sqrs := listMapper(xi, func(i int, v int) int {
// 		return v * v
// 	})

// 	fmt.Println(sqrs)

// 	fmt.Print("\n\n\n\n")

// 	b := bar(12)

// 	fmt.Println(b(1))
// 	fmt.Println(b(2))
// 	fmt.Println(b(3))

// }

// func listMapper(vals []int, cb func(i int, v int) int) []int {
// 	var results []int
// 	for i, v := range vals {
// 		results = append(results, cb(i, v))
// 	}
// 	return results
// }

// func foo(init int) func(x int) int {
// 	return func(x int) int {
// 		init = init * x
// 		return init
// 	}
// }

// func bar(init int) func(x int) int {
// 	return func(x int) int {
// 		return init * x
// 	}
// }

// // func a(x int) {
// // 	fmt.Print("int\n")
// // }

// // func a(x string) {
// // 	fmt.Print("string \n")
// // }
