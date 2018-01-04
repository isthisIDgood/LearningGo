package main

import "fmt"

func main() {
	n := greatestInt(56, 79, 88, 9001, 852, 167)
	fmt.Println(n)
}

func greatestInt(i ...int) int {
	max := 0
	for _, v := range i {
		if max < v {
			max = v
		}
	}
	return max
}
