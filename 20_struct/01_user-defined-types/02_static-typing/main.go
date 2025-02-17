package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 20
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 21
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work:
	//	 fmt.Println(myAge + yourAge)

	// this conversion works:
	//	 fmt.Println(int(myAge) + yourAge)
}
