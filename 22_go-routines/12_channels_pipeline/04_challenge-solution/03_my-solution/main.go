package main

import "fmt"

func main() {
	n := numbers()
	f := factorialChan(n)
	for x := range f {
		fmt.Println(x)
	}
}

func numbers() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 9; i++ {
			for j := 1; j <= 10; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorialChan(n chan int) chan int {
	out := make(chan int)
	go func() {
		for x := range n {
			out <- factorialInt(x)
		}
		close(out)
	}()
	return out
}

func factorialInt(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
