package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Arch\t", runtime.GOARCH)
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("CPU\t", runtime.NumCPU())

	fmt.Println(reflect.TypeOf(wg))
	fmt.Printf("%+v\n", wg)
	wg.Add(1)
	fmt.Println(wg)

	go foo()
	fmt.Println(wg)

	bar()
	fmt.Println(wg)

	wg.Wait()
	fmt.Println(wg)

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
