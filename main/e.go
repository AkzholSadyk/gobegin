package main

import "fmt"

func getFour(a int) int {
	if a != 4 {
		return 4
	}
	return 0
}

func main() {
	a := 5

	switch a {
	case 1:
		fmt.Println("Monday")
	case 5:
		fmt.Println("Friday")
	case 2, 3:
		fmt.Println("Tuesday or Wednesday")
	case getFour(a):
		fmt.Println("Four")

	default:
		fmt.Println("Not a weekday")
	}
}
