package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	writeWithMutex()
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final counter:", counter)
	fmt.Println("Elapsed time:", time.Since(start).Seconds())
}
