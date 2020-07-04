package atomicds

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var incrementor int64

	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		fmt.Println("goroutine number\t", i)

		go func(i int) {
			atomic.AddInt64(&incrementor, 1)
			// fmt.Println("i", incrementor)
			v := atomic.LoadInt64(&incrementor)
			fmt.Println(i, v)
			wg.Done()
		}(i)

		fmt.Println("NUM GOROUTINE\t\t", runtime.NumGoroutine())
	}
	fmt.Println("waitng")
	wg.Wait()
	fmt.Println("Counter", incrementor)

}
