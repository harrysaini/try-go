package starting_code

import (
	"fmt"
	"harish/try-go/exercises/011-testing/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(4),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
