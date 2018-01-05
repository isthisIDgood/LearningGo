package main

import "fmt"

func main() {
	var x [256]int
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}
