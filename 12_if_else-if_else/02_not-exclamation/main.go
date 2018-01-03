package main

import (
	"fmt"
)

func main(){
	if !true{
		fmt.Println("This won't run")
	}

	if !false{
		fmt.Println("This will run")
	}
}
