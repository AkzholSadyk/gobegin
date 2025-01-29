package main

import "fmt"

func main() {
	a := 1

	switch {
	case a == 1:
		fmt.Println("Monday")
		a = 5
		fallthrough // switch ти токтатып тастамай жалгастыра береди

	case a == 5:
		fmt.Println("Friday")
	case a == 2 || a == 3:
		fmt.Println("Tuesday or Wednesday")

	default:
		fmt.Println("Not a weekday")
	}
}
