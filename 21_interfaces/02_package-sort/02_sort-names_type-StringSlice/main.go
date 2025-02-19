package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Dimas", "ADI", "Sarah", "Jenny"}
	fmt.Println(s)
	//	sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

}

// https://golang.org/pkg/sort/#Sort
