package main

import "fmt"

func main() {
	fmt.Println(greet("DImas ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
