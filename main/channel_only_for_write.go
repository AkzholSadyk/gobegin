package main

import (
	"fmt"
	"time"
)

func main() {
	channel_only_for_write()
}

func channel_only_for_write() {
	unBufferedChannel := make(chan<- int) //вот здесь они отличаются этот канал только для записи
	fmt.Printf("Len: %d Cap: %d\n", len(unBufferedChannel), cap(unBufferedChannel))

	// blocks until smb reads
	//unBufferedChannel <- 1

	//  blocks on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		//chanForWriting <- 1 // Теперь используем переданный канал
		unBufferedChannel <- 1
	}(unBufferedChannel)

	//value := <-unBufferedChannel // Читаем значение из канала
	//fmt.Println(value)

	// blocks on writing then read
	//go func(chanForReading <-chan int) {
	//	time.Sleep(time.Second)
	//	value := <-chanForReading
	//	fmt.Println(value)
	//}(unBufferedChannel)

	unBufferedChannel <- 3

	go func() {
		time.Sleep(time.Millisecond * 500)
		close(unBufferedChannel)
	}()

	go func() {
		time.Sleep(time.Second)
		//fmt.Println(<-unBufferedChannel)
	}()
	unBufferedChannel <- 4
}
