package main

import "fmt"

type OurString string
type OurInt int

func main() {
	var customString OurString = "Hello"
	var customInt OurInt = 5
	fmt.Printf("%T %v %T %v", customString, customString, customInt, customInt)

	customInt = OurInt(5)
}
