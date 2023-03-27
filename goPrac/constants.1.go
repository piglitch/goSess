package main
import ("fmt")

const(
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

func main(){
	var specialstType int
	fmt.Printf("%v\n", specialstType == catSpecialist)
	fmt.Println(dogSpecialist)
}

