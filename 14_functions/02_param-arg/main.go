package main

import "fmt"

func main() {
	greet("dimas")
	greet("sarah")
}

func greet(name string) {
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
