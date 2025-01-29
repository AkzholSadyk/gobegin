package main

import "fmt"

var erasyl bool = true

const ersultan = "elom"

func main() {
	eki := 2
	bir := 1

	akzhol := "sadyk"
	var ako string = "akooo"
	ako = "akoosh"
	fmt.Printf("Type: %T  Value: %v\n", ako, ako)
	fmt.Printf("Type: %T  Value: %v\n", akzhol, akzhol)

	var erasyl bool = false

	fmt.Println(erasyl)

	fmt.Println(ersultan)

	hello := fmt.Sprintf("hello %s", ako)
	_ = hello
	fmt.Println(hello)

	foo()

	foo2(ako)

	eba := foo3(ako, akzhol)
	fmt.Println(eba)

	ekiznak(eki, bir)
}
func multiply(eki int, bir int, ako string) (aty string, sum int) {
	sum = eki * bir
	aty = fmt.Sprintf("%s %d", ako, sum)
	return
}

func ekiznak(eki, bir int) (int, int) {
	return eki + bir, eki * bir
}

func foo3(ako, akzhol string) string {
	return ako + akzhol
}

func foo2(ako string) {
	fmt.Printf("foo2 %s\n", ako)
}

func foo() {
	fmt.Println("foo")
}
