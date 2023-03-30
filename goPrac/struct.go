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
			"Pace", "Dribbling", "Poweer shot",
		},
	}
	fmt.Println(aPlayer.attributes)
}