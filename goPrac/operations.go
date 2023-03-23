package main

import ("fmt")

func main(){
	var num, num1 int
	var ops string

	fmt.Print("Type first number: ")	
	fmt.Scan(&num) // To take input 
	fmt.Print("Type second number: ")
	fmt.Scan(&num1)
	fmt.Println("1. Sum\n2. Sub\n3. Mult\n4. Div\n5. Remainder")
	fmt.Print("Choose one from 1, 2, 3, 4, 5: ")
	fmt.Scan(&ops)


	switch ops{
	case "1":
		fmt.Println("Value after multiplication: ", num+num1)
	case "2":
		fmt.Println("Value after multiplication: ", num-num1)
	case "3":
		fmt.Println("Value after multiplication: ", num*num1)
	case "4":
		fmt.Println("Value after multiplication: ", num/num1)
	case "5":
		fmt.Println("Value after multiplication: ", num%num1)			
	
	default:
		fmt.Println("Wrong input! Please choose from 1, 2, 3, 4, 5")	
	}
	

	
	
}
