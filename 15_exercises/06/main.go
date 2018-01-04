package main

import "fmt"

func main() {
	x := 1
	y := 2
	sum := 0
	for y < 4000000 {
		if y%2 == 0 {
			sum += y
		}
		temp := y
		y = x + y
		x = temp
	}
	fmt.Println(sum)
}