package arrays
import ("fmt")
// func main(){
// 	var students [3]string
// 	students[2] = "Antony"
// 	fmt.Printf("Students: %v\n", students)
// 	arrayOfArrays()
// 	slice()
// 	slack()
// }

func ArrayOfArrays(){
	var randomMatrix [2][2]int = [2][2]int{[2]int{1, 0}, [2]int{0,1}}
	fmt.Println(randomMatrix)
	fmt.Println(len(randomMatrix))
	fmt.Println(cap(randomMatrix))
}

func Slice(){
	a := []int{1,2,3,4,5,6,7,8,9}
	b := a[:] //slice of all
	c := a[5:] // from the 6th element to the end
	d := a[:8] // upto 8th element
	e := a[4:7] //5th to 7h element
	fmt.Println(a,b,c,d,e) 
}

func Slack(){
	a:= []int{1,2,3,4,6,7,8,9}
	fmt.Println(a)
	b:= append(a[:3], a[4:]...)			// To remove 4 from thee array
	fmt.Println(b)
}
