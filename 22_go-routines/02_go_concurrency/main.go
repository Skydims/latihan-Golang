package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 24; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 24; i++ {
		fmt.Println("Bar:", i)
	}
}
