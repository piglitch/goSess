package main
import (
	"fmt"
	"mymodule/arrays"
)
func main(){
	var students [3]string
	students[2] = "Antony"
	fmt.Printf("Students: %v\n", students)
	arrays.ArrayOfArrays()
	arrays.Slice()
	arrays.Slack()
}