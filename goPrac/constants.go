package main
import ("fmt")
const myConst int = 24	// Declared globally

func main(){
	const myConst int = 42 
	
	fmt.Printf("%v, %T\n", myConst, myConst) // Returns the value declared inside the function
	anotherConst() // Call a function that returns the value of that was globally declared
}
func anotherConst(){
	fmt.Printf("%v, %T\n", myConst, myConst)
}