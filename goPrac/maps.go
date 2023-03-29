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
	nameAge["Bruno"] = 28
	delete(nameAge, "Bruno")
	pop, ok := nameAge["Sancho"] //Using ok we can see if the value we popped does or does not belong to the map.
	fmt.Println(pop, ok)
}