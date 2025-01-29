package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 2
	fmt.Printf("Type: %T  Value: %v\n", slice, slice)
	fmt.Printf("Len: %d  Cap: %d\n\n", len(slice), cap(slice))

	withAppend := append(slice[:i], slice[i+1:]...)
	fmt.Printf("Type: %T  Value: %v\n", withAppend, withAppend)
}
