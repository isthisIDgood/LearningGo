package main

import (
	"fmt"
)

func main() {

	integer := func(x int) {
		fmt.Println(x / 2)
		if x%2 == 0 {
			fmt.Println(x, true)
		} else {
			fmt.Println(x, false)
		}
	}
	integer(56)
}
