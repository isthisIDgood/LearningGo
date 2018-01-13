package main
import (
	"fmt"
	"net/http"
)
func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go doPrint(v, done)
	}
	// wait for all goroutines to complete before exiting
	for  range values {
		<-done
	}
}
func doPrint(v string, done chan bool) {
	http.Get("https://www.google.com/")
	fmt.Println(v)
	done <- true
}