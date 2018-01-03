package main

import (
	"fmt"
)

func main(){
	var number int
	var largeNumber int
	fmt.Print("Please enter a small number ")
	fmt.Scan(&number)
	fmt.Print("Please enter a large number ")
	fmt.Scan(&largeNumber)
	fmt.Println("if you divide your larger number by your smaller number you have a remainder of", largeNumber % number)

}
