package main

import "fmt"

func main() {
	a := 10
	c := 0
	for b := 0; b < a; b++ { //mynau forik
		if c == 4 {
			break
		}
		c++
		fmt.Printf("%d ", a)
	}
	fmt.Println("\n")
	for c < a { //mynau while
		c++
		fmt.Printf("%d ", a)
	}
}
