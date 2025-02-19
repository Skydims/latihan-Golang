package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Dimas", "John", "Al", "Jenny"}
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

// https://golang.org/pkg/sort/#Sort
