package main

import "fmt"

func main() {
	a := 7
	b := 7
	if a > b {
		fmt.Println(a)
	} else if a < b {
		fmt.Println(b)
	} else {
		fmt.Println("a == b")
	}

}
