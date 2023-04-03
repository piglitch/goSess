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
	nest()
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


func nest(){
Loop:	
	for i := 1; i <= 3; i++{
		for j:=1; j <= 3; j++{
			fmt.Println(i*j)
			if i*j > 3{
				break Loop
			}
		}
	}
}