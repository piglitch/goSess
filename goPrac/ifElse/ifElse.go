package main
import (
	"fmt"
)

func main(){
	var guess int
	var cont string
	number:= 50
	fmt.Print("Take a guess: ")
	fmt.Scan(&guess)

	if number > guess  {
		fmt.Println("Too low")		
	}	
	if number == guess {
		fmt.Println("You got it!")
	}
	if number < guess {
		fmt.Println("Too high")
	}
	fmt.Print("Continue? y/n: ")
	fmt.Scan(&cont)
	if cont == "y" {
		main()		
	}	else{	fmt.Println("Thanks!")
		
	}
	typeSwitch()

}

func typeSwitch(){
	var i interface{} = [3]int{}
	switch i.(type){
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("array of 3 integers")
	default:
		fmt.Println("i is another type")	 			
	}
}