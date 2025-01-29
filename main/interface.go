package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

type Duck struct {
	Name, Surname string
}

func (h Human) Run() string {
	return fmt.Sprintf("%s is running", h.Name)
}

func (d Duck) Run() string {
	return fmt.Sprintf("%s %s is waddling instead of running", d.Name, d.Surname)
}

func interfaceValues() {
	var runner Runner
	fmt.Printf("%T %v\n", runner, runner)

	if runner == nil {
		fmt.Println("nil")
	}

	runner = Human{"John"}
	fmt.Println(runner.Run())
	fmt.Printf("%T %v\n", runner, runner)
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("%T %v\n", runner, runner)

	// Объект John выполняет метод Run
	john := &Human{"John"}
	polymorphism(john)

	// Объект Donald Duck выполняет метод Run
	donald := &Duck{"Donald", "Duck"}
	polymorphism(donald)

}

func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

func main() {
	interfaceValues()

	typeAssertionAndPolymorphism()
}
