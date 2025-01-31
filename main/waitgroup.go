package main

import (
	"fmt"
	"sync"
)

func main() {

	//withoutWait()

	withWait()

}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			//defer wg.Done()
			fmt.Println(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}
