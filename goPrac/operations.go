package main

import ("fmt")

func main(){
	var num, num1 int

	fmt.Print("Type first number: ")	
	fmt.Scan(&num) // To take input 
	fmt.Print("Type second number: ")
	fmt.Scan(&num1)

	var mult  int = num*num1

	fmt.Println("Value after multiplication: ", mult)
	
}
