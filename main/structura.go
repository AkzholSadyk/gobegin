package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var John Person
	fmt.Printf("%T %v \n	", John, John)

	John = Person{"John", 25}
	fmt.Printf("%T %v \n	", John, John)

	John.Age = 26
	John.Name = "Akzhol"
	fmt.Println(John)

	Brad := Person{"Brad", 27}
	fmt.Println(Brad)

	pBrad := &Brad
	fmt.Println((*pBrad).Name)
	fmt.Println(pBrad.Name)

	pIvan := &Person{"Ivan", 90}
	fmt.Println(pIvan.Name)
	fmt.Println(pIvan)

	unnamed := struct {
		Name, LastName, BirthDate string
		Age                       int
	}{
		Name:      "Ivan",
		LastName:  "Ivanov",
		BirthDate: "1990-01-01",
		Age:       90,
	}
	fmt.Println(unnamed)
}
