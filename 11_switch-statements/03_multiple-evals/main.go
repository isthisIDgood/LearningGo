package main

import "fmt"

func main(){
	switch "Jenny"{
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with m")
	case "Julian", "Susy":
		fmt.Println("Wassup Julian/Susy")
	}
}
