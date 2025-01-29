package main

import "fmt"

func square(num int) {
	num = num * num
}

func squarePointer(num *int) {
	*num *= *num
}

func hasWallet(money *int) bool {
	return money != nil
}

func main() {
	num := 3
	square(num)
	println(num)

	squarePointer(&num)
	fmt.Println(num)

	wallet := 0
	fmt.Println(hasWallet(&wallet))

	wallet1 := 100
	fmt.Println(hasWallet(&wallet1))

	var wallet2 *int
	fmt.Println(hasWallet(wallet2))

}
