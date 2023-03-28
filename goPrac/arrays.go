package main
import ("fmt")

func main(){
	var students [3]string
	students[2] = "Antony"
	fmt.Printf("Students: %v\n", students)
	arrayOfArrays()
}

func arrayOfArrays(){
	var randomMatrix [2][2]int = [2][2]int{[2]int{1, 0}, [2]int{0,1}}
	fmt.Println(randomMatrix)
}

