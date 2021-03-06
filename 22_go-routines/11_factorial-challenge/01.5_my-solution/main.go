package main

import (
	"fmt"
)

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
			out <- total
		}
		close(out)
	}()
	return out
}
//	go func() {
//		total := 1
//		for i := n; i > 0; i-- {
//			total *= i
//		}
//		c <- total
//		close(c)
//	}()
//	return c
//}

//takes original answer and puts it into the channel twice
