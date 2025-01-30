package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())

	go ako(100)

	//runtime.Gosched() //вручную переключение горутины
	time.Sleep(time.Second)
	fmt.Println("exit")
}

func ako(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}
