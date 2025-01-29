package main

import "fmt"

func main() {
	var ako []int
	ako = append(ako, 1, 2, 3, 4, 6)

	for i := 0; i < len(ako); i++ {
		fmt.Println(ako[i])
	}

	fmt.Println(ako)
	stringSliceLiteral := []string{"a", "b", "c"}
	fmt.Println(stringSliceLiteral)
	fmt.Printf(("len: %d cap: %d"), len(stringSliceLiteral), cap(stringSliceLiteral))

	sliceByMake := make([]int, 5, 10)
	fmt.Println(sliceByMake)
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Println(sliceByMake)
}
