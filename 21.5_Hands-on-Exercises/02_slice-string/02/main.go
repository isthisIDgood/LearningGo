package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	//Reverses type StringSlice and returns an interface which gets passed into sort.Sort
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	// works to sort type StringSlice
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
