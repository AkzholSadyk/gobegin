package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	people := [2]Person2{
		{
			Name: "John",
			Age:  25,
		},
		{
			Name: "Brad",
			Age:  27,
		},
	}

	for i := 0; i < len(people); i++ {
		fmt.Printf("%v %v\n", people[i].Name, people[i].Age)
	}

}
