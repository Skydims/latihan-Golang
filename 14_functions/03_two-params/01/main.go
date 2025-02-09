package main

import "fmt"

func main() {
	greet("Dimas", "Sarah")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
