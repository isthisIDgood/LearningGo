package main

import (
	"fmt"
	"net/http"
)

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(n string) {
			http.Get("https://www.google.com/")
			fmt.Println(n)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}

/*
Even easier is just to create a new variable,
using a declaration style that may seem odd but works fine in Go.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
