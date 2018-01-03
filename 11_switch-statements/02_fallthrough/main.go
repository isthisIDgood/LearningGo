package main

import "fmt"

func main(){
	switch "Marcus"{
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("wassup Julian")
	case "Susy":
		fmt.Println("wassup Susy")
	}
}
