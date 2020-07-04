package inc

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incrementor := 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		fmt.Println("goroutine number\t", i)

		go func() {
			v := incrementor
			runtime.Gosched()
			fmt.Println(v)
			v++
			incrementor = v
			wg.Done()
		}()

		fmt.Println("NUM GOROUTINE\t\t", runtime.NumGoroutine())
	}
	fmt.Println("waitng")
	wg.Wait()
	fmt.Println("Counter", incrementor)

}
