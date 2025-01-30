package main

import "fmt"

func main() {
	//defer block //ожидание
	defer fmt.Println(2)
	defer fmt.Println(1)

	fmt.Println("exit")

	fmt.Println(sum(2, 5))
}

func sum(a, b int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = a + b
	return
}
