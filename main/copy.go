package main

import "fmt"

func main() {
	destination := make([]string, 0, 2)
	source := []string{"Vasya", "Petya", "Katya"}

	fmt.Println("Copied: ", copy(destination, source))
	fmt.Printf("Type: %T Value: %v\n", destination, destination)
	fmt.Printf("length: %d Capacity: %d\n\n", len(destination), len(destination))

	destination = make([]string, 2, 3)
	fmt.Println("Copied: ", copy(destination, source))
	fmt.Printf("Type: %T Value: %v\n", destination, destination)
	fmt.Printf("length: %d Capacity: %d\n\n", len(destination), len(destination))

	destination = make([]string, len(source))
	fmt.Println("Copied: ", copy(destination, source))
	fmt.Printf("Type: %T Value: %v\n", destination, destination)
	fmt.Printf("length: %d Capacity: %d\n\n", len(destination), len(destination))

}
