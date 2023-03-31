package main
import (
	"fmt"
)

func main(){
	var guess int
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
}