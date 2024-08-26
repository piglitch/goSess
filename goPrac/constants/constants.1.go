package main
import ("fmt")

const myConst int = 42
const(
	_ = iota
	KB = 1 << (10*iota)
	MB
	GB
)

const(
	isAdmin = 1 << iota
	isHeadquarters 
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeNorthAmerica
	canSeeSouthAmerica
	canSeeEurope
)

func main(){

	fileSize := 4000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)
	fmt.Println(KB,GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeNorthAmerica		// bitwise OR "|"
	fmt.Printf("%b\n", roles)
	fmt.Printf("%v", isAdmin & roles == isAdmin)		// bitwise AND "&"


	const myConst int = 30 
	
	fmt.Printf("%v, %T\n", myConst, myConst) // Returns the value declared inside the function
	anotherConst() // Call a function that returns the value of that was globally declared
}
func anotherConst(){
	fmt.Printf("%v, %T\n", myConst, myConst)
}