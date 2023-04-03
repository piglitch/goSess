package main

import (
	"fmt"
)

func main(){
	for i:=0; i<10; i++{
		if i%2 != 0{	//To only print even numbers
			continue
		}
		fmt.Println(i)
	}
	doWh()
}

// Using if statement to mimic do while loop

func doWh(){
	i:=0
	for{
		fmt.Println(i)
		i++
		if i == 7 {
			break			
		}
	}
	
}