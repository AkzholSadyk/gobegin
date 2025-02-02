package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//baseKnowledge()
	workerpool() // часта колданылатын паттерн
}

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	toDo := context.TODO()
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err()) // context те ошибка барма деген соз
	cancel()
	fmt.Println(withCancel.Err())

	fmt.Printf("\n")
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	fmt.Printf("\n")

	withTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	fmt.Println(<-withTimeout.Done())
}

func workerpool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcees, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcees, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			if i == 500 {
				cancel()
			}
			numbersToProcees <- i
		}
		close(numbersToProcees)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for result := range processedNumbers {
		counter++
		fmt.Println(result)
	}
	fmt.Println(counter)
}

func worker(ctx context.Context, ToProcees, processed chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-ToProcees:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value

		}
	}
}
