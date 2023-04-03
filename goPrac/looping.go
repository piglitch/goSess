package main

import (
	"fmt"
)

func main(){
	for i:=0; i<5; i++{
		fmt.Println("name's bond")
	}
	doWh()
}

// Using if statement to mimic do while loop

func doWh(){
	i:=0
	for{
		fmt.Println("Name's bond")
		i++
		if i == 7 {
			break			
		}
	}
	
}