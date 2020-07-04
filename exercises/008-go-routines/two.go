package main

import (
	"fmt"
	"math/rand"
)

type person struct {
	name string
	age  int
}

/*
	Main function
*/
func (p *person) speak() {
	fmt.Printf("I am %v of age %v\n", p.name, p.age)
}

type human interface {
	speak()
}

//
// 	Main function
//
func saySomething(h human) {
	h.speak()
}

var x int = randomNumber()

func randomNumber() int {
	return rand.Int()
}

func main() {

	fmt.Println("x", x)

	p1 := person{
		name: "Ram",
		age:  12,
	}

	saySomething(&p1)
	p1.speak()
}
