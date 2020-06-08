// package main

// import (
// 	"fmt"
// 	"math"
// )

// type square struct {
// 	length float64
// }

// type circle struct {
// 	radius float64
// }

// func (s square) area() float64 {
// 	return s.length * s.length
// }

// func (c circle) area() float64 {
// 	return math.Pi * (c.radius * c.radius)
// }

// type shape interface {
// 	area() float64
// }

// func info(s shape) {
// 	fmt.Printf("%T\n", s)
// 	fmt.Println(s.area())
// }

// func main() {
// 	fmt.Println(22.0 / 7)
// 	fmt.Println(math.Pi)

// 	c := circle{
// 		radius: 7,
// 	}

// 	info(c)

// 	s := square{
// 		length: 12,
// 	}

// 	info(s)
// }
