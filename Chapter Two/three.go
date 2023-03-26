package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type SaiyanThree struct {
	*Person
	Power int
}

func main() {
	goku := &SaiyanThree{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
}
