package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)

	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan<- string) {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- fmt.Sprintf("i-> %d, j->%d", i, j)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(c)

}

func receive(c <-chan string) {
	for i := range c {
		fmt.Println(i)
	}
}
