package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("--------")

	j := 0
	for {
		j++
		if j > 10 {
			break
		}
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)

	}

	fmt.Println("--------")

	for k := 33; k <= 122; k++ {
		fmt.Printf("%v\t%#x\t%#U\n", k, k, k)
	}

	fmt.Println("----NESTED----")

	for i := 0; i <= 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		for j := 10 - i; j >= 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
