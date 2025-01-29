package main

import "fmt"

func main() {
	var intPointer *int
	fmt.Printf("%T %v \n", intPointer, intPointer)

	a := 7
	var intPointer2 *int = &a
	fmt.Printf("%T %v %v \n", intPointer2, intPointer2, *intPointer2)

	var newPointer = new(float64)
	fmt.Printf("%T %v %v\n", newPointer, newPointer, *newPointer)
	*newPointer = 7.7
	fmt.Printf("%T %v %v\n", newPointer, newPointer, *newPointer)
}
