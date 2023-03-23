package main
import (
	"fmt"
)
func main(){
	n:= 1 == 1
	m:= 1 == 2
	fmt.Printf("%v, %T\n", n, n) //Prints true because 1 == 1
	fmt.Printf("%v, %T\n", m, m)	//Prints false bacause 1 is not equal to 2
	bitshift()
} 

func bitshift(){
	a:=8
	fmt.Println(a<<3)
	fmt.Println(a>>3)
}

