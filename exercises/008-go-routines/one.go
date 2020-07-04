package tets

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	fmt.Printf("\n%+v\n", wg)

	go foo()
	go bar()

	fmt.Print("\nWaitng\n")
	wg.Wait()

	fmt.Print("\nDone\n")
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Print(`foo`, i)
	}
	wg.Done()
	fmt.Printf("\n%+v\n", wg)
}

func bar() {
	for i := 0; i < 10000; i++ {
		fmt.Print(`bar`, i)
	}
	wg.Done()
}
