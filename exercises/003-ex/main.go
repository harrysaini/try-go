package main

import "fmt"

func main() {
	fmt.Println("---> 1")

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	fmt.Println("---> 2")

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U \n", i)
		}
	}

	fmt.Println("---> 3")

	y := 1996

	for y <= 2020 {
		fmt.Println(y)
		y++
	}

	fmt.Println("---> 4")

	l := 1996

	for {
		if l > 2020 {
			break
		}
		fmt.Println(l)
		l++
	}

	fmt.Println("---> 5")

	for i := 10; i <= 100; i++ {
		fmt.Printf("%v%%4 = %v \n", i, i%4)
	}

	fmt.Println("---> 6,7")

	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even \n", i)
		} else if i%3 == 0 {
			fmt.Printf("%v is divisible by 3 not by 2 \n", i)
		} else {
			fmt.Printf("%v is not divisible by 3 and 2 \n", i)
		}
	}

	fmt.Println("---> 8")

	switch {
	case 1 == 1:
		fmt.Println("1==1")
	case 1 != 2:
		fmt.Println("1!=2")
	}

	fmt.Println("---> 9")

	favSport := "cricdket"

	switch favSport {
	case "cricket":
		fmt.Println("cricket is fun")
	case "footbalss":
		fmt.Println("footlbaal is fun")
	default:
		fmt.Println("Default")
		fmt.Println(favSport)

	}

	fmt.Println("---> 10")

	fmt.Println(1 == 1 && true)
	fmt.Println(true && false)
	fmt.Println(1 == 1 || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
