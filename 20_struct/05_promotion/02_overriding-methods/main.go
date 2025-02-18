package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("aku adalah cucu ultramen.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("apan tuh man.")
}

func main() {
	p1 := person{
		Name: "Dimas adi saputra",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
