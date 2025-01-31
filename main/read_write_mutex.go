package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	readWithMutex()
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			//_ = counter

			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
	fmt.Println("Elapsed time:", time.Since(start).Seconds())
}
