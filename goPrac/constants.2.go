package main
import(
	"fmt"
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
	var roles byte = isAdmin | canSeeFinancials | canSeeNorthAmerica		// bitwise OR "|"
	fmt.Printf("%b\n", roles)
	fmt.Printf("%v", isAdmin & roles == isAdmin)		// bitwise AND "&"
}