package main

import "fmt"

func main() {
	intArr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T %v", intArr, intArr)

	intSlice := intArr[1:4]
	fmt.Printf("%T %v", intSlice, intSlice)

	intSlicewq2 := intArr[:]
	fmt.Printf("%T %v", intSlicewq2, intSlicewq2)

}
