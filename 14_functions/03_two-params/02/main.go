package main

import "fmt"

func main() {
	greet("Dimas", "Doe")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
