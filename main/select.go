package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	baseSelect()
	//gracefulShutdown()
}

func baseSelect() {
	bufferedChannel3 := make(chan string, 1)
	bufferedChannel3 <- "first"

	select {
	case str := <-bufferedChannel3:
		fmt.Println("read", str)
	case bufferedChannel3 <- "second":
		fmt.Println("write", <-bufferedChannel3, <-bufferedChannel3)
	}

	unbufChan := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		unbufChan <- 1
	}()

	select {
	//case bufferedChannel3 <- "Third":
	//	fmt.Println("unblocking writing")
	case val := <-unbufChan:
		fmt.Println("blocking writing", val)
	case <-time.After(time.Millisecond * 2500):
		fmt.Println("timed out")
		//default:
		//	fmt.Println("nothing to do")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second)

	go func() {
		defer close(resultChan)

		for i := 0; i < 1000; i++ {

			select {
			case <-timer:
				fmt.Println("times up")
				return

			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	for {
		select {
		case <-timer:
			fmt.Println("timed out")
			return
		case sig := <-sigChan:
			fmt.Println("stopped by signal:", sig)
			return
		}
	}
}
