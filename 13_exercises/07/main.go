package main

import "fmt"

func main (){
	var sumAnswer = 0
	for i := 0; i < 1000; i++{
		if i % 3 == 0{
			sumAnswer += i
		}else if i % 5 == 0{
			sumAnswer += i
		}
	}
	fmt.Println(sumAnswer)
}
