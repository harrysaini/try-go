// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type person struct {
// 	firstName string
// 	lastName  string
// }

// func (p *person) speak() {
// 	fmt.Println("speaxk", p)
// 	fmt.Printf("I am %v %v\n", p.firstName, p.lastName)
// }

// func main() {
// 	x := 10

// 	fmt.Println(x)
// 	fmt.Println(&x)
// 	fmt.Println(reflect.TypeOf(x))

// 	fmt.Println("--->")

// 	p := person{
// 		firstName: "Ram",
// 		lastName:  "Dev",
// 	}

// 	fmt.Println(p)
// 	p.speak()
// 	changeMe(&p)
// 	fmt.Println(p)

// 	fmt.Println(reflect.TypeOf("sfgdfg"))
// }

// func changeMe(p *person) {
// 	ps := *p
// 	p.speak()
// 	ps.speak()
// 	(*p).speak()
// 	ps.firstName = "Kam"
// 	p.firstName = "KAMU"
// }
