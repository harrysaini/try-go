package main

import "fmt"

type person struct {
	firstName          string
	lastName           string
	favouriteIceCreams []string
}

func main() {
	p1 := person{
		firstName:          "ram",
		lastName:           "dev",
		favouriteIceCreams: []string{"vanilla", "chochlate"},
	}

	p2 := person{
		firstName:          "kam",
		lastName:           "devu",
		favouriteIceCreams: []string{"hot chochlate", "strawberry"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favouriteIceCreams {
		fmt.Println(i, v)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println("--------->", m)

	for _, v := range m {
		fmt.Println("---", v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favouriteIceCreams {
			fmt.Println(i, val)
		}
	}

	embeddedStructs()

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type pedan struct {
	v      vehicle
	luxury bool
}

func embeddedStructs() {
	fmt.Println(`
--->
	`)

	t := truck{
		vehicle: vehicle{
			doors: 12,
			color: `red`,
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: `dark blue`,
		},
		luxury: false,
	}

	fmt.Println(t.vehicle, t.doors, t.fourWheel, t.color)
	fmt.Println(s.vehicle, s.doors, s.luxury, s.color)

	fmt.Printf("%T\n", t.vehicle)

	fmt.Printf("%T\n", s)

	sp := pedan{
		luxury: false,
	}

	fmt.Printf("abc %T\t%v\n", sp, sp)
	fmt.Println(sp.v)

	anonymousStruct(67)
}

func anonymousStruct(v int) {
	fmt.Println(`
--- anaonymous->
	`)

	s := struct {
		name string
	}{
		name: "abc",
	}

	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
