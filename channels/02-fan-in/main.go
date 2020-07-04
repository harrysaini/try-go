package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe", false), boring("Ann", true))
	for i := 0; i < 1000; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string, moreWait bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			var gap int
			if gap = 1; moreWait {
				gap = 10
			}

			time.Sleep(time.Duration(rand.Intn(1e3)*gap) * time.Millisecond)

		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
