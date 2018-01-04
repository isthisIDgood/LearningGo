package main

import "fmt"

func main() {
	x := 1
	y := 2
	var sum int
	for y < 4000000 {
		if y%2 == 0 {
			sum += y
		}
		x, y = y, x+y
	}
	fmt.Println(sum)
}

// prints the sum of the even numbers in the fibonacci sequence whose values don't equal or surpass 4 million