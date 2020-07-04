package main

import (
	"fmt"
	"sync"
)

func main() {
	incrementor := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)

	for i := 0; i < 100; i++ {
		// fmt.Println("goroutine number\t", i)

		go func(i int) {
			mu.Lock()
			v := incrementor
			v++
			incrementor = v
			fmt.Println("i", i, "v", v)
			mu.Unlock()
			wg.Done()
		}(i)

		// fmt.Println("NUM GOROUTINE\t\t", runtime.NumGoroutine())
	}
	fmt.Println("waitng")
	wg.Wait()
	fmt.Println("Counter", incrementor)

}
