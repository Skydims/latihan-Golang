package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main() {
	p1 := person{"Dimas", "Adi", 20}
	p2 := person{"Sarah", "julian sari", 21}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
