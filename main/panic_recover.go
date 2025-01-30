package main

import "fmt"

func main() {

	fmt.Println("exit")
	makePanic()
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()
	panic("some PANIC")
}
