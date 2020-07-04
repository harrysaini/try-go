package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type vehicle struct {
	Car string
}

type person struct {
	First   string
	Last    string
	Sayings []string
	V       vehicle
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))

	m := map[string]int{
		"CAR": 1,
	}

	bs, err = json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))

}
