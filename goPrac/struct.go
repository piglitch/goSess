package main
import(
	"fmt"
)

type player struct{
	name string
	position string
	age int
	attributes []string
}

func main(){
	aPlayer:= player{
		name: "Rashford",
		position: "LW",
		age: 25,
		attributes: []string{
			"Pace", "Dribbling", "Power shot",
		},
	}
	fmt.Println(aPlayer.attributes)
	shortStruct()
}

func shortStruct(){				// A fast and short way of defining a struct (not recommended for most cases)
	myself:= struct{name string}{name: "Avi"}
	myName:= &myself			// Use '&' to change the value of myself and assign it to myName. myName is a pointer
	myName.name = "Graham"		// that points to the address of mySelf
	fmt.Println(myself)
	fmt.Println(myName)
}