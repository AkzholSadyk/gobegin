package main

import "fmt"

func main() {
	var multipleir func(x, y int) int
	multipleir = func(x, y int) int { return x * y }
	fmt.Println(multipleir(2, 3))
	devider := func(x, y int) int { return x / y }
	fmt.Println(devider(9, 3))
}
