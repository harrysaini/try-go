package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and my age is %v \n", p.firstName, p.lastName, p.age)
}

// type human interface {
// 	speak()
// }

type superPerson struct {
	person
	superPower string
}

// func dance(p person) {
// 	p.firstName = "Abc"
// 	fmt.Println(p)
// }

func main() {
	p1 := person{
		firstName: "Ram",
		lastName:  "Dev",
		age:       43,
	}
	p1.speak()

	p2 := person{
		firstName: "Krishan",
		lastName:  "Saini",
		age:       21,
	}

	p2.speak()

	fmt.Println(p1)
	// dance(p1)
	fmt.Println(p1)

	sp1 := superPerson{
		person: person{
			firstName: "Krishan",
			lastName:  "Saini",
			age:       21,
		},
		superPower: "dancer",
	}

	sp1.speak()
	fmt.Println(sp1)
}
