package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 1; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

// % gives the remainder of dividing the two numbers (13 / 3 gives remainder 1)
