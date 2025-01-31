package main

import (
	"fmt"
	"time"
)

func main() {
	//nilChannel()
	//unBufferedChannel()
	//bufferedChannel()
	forRange()
}

func nilChannel() {
	var nilChannel chan int
	fmt.Printf("Len: %d  Cap: %d\n", len(nilChannel), cap(nilChannel))

	//write to nil channel blocks forever
	//nilChannel <- 1

	//read from nil channel blocks forever
	//<-nilChannel

	//closing nil channel will raise a panic
	//close(nilChannel)
}

func unBufferedChannel() {
	unBufferedChannel := make(chan int) //этот канал двухнапрвленний можно и читать и писать
	fmt.Printf("Len: %d Cap: %d\n", len(unBufferedChannel), cap(unBufferedChannel))

	// blocks until smb reads
	//unBufferedChannel <- 1

	//  blocks on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		//chanForWriting <- 1 // Теперь используем переданный канал
		unBufferedChannel <- 1
	}(unBufferedChannel)

	value := <-unBufferedChannel // Читаем значение из канала
	fmt.Println(value)

	// blocks on writing then read
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unBufferedChannel)

	unBufferedChannel <- 3

	go func() {
		time.Sleep(time.Millisecond * 500)
		close(unBufferedChannel)
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-unBufferedChannel)
	}()
	unBufferedChannel <- 4
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

}

func forRange() {
	bufferedChannel := make(chan int, 3)

	numbers := []int{1, 2, 3, 4, 5}

	//show all elements with for
	go func() {
		for _, number := range numbers {
			bufferedChannel <- number
		}
		close(bufferedChannel)
	}()

	for {
		//v := <-bufferedChannel
		//fmt.Println(v)
		value, ok := <-bufferedChannel
		fmt.Println(value, ok)
		if !ok {
			break
		}
	}

	//show with for range buffered
	bufferedChannel2 := make(chan int, 3)

	go func() {
		for _, number := range numbers {
			bufferedChannel2 <- number
		}
		close(bufferedChannel2)
	}()

	for value := range bufferedChannel2 {
		fmt.Println("buffered: ", value)
	}

	//unbuffered
	unBufferedChannel2 := make(chan int)
	go func() {
		for _, number := range numbers {
			unBufferedChannel2 <- number
		}
		close(unBufferedChannel2)
	}()

	for value := range unBufferedChannel2 {
		fmt.Println("unBuffered: ", value)
	}
}
