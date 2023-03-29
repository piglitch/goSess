package main
import (
	"fmt"
)

func main(){
	nameAge := make(map[string]int)
	nameAge = map[string]int{
		"Rashy": 25,
		"Martial": 27,
		"Garnacho": 18,
		"Sancho": 23,
	}
	fmt.Println(nameAge["Martial"])
}