package main	// main package is made of the files that are in the same directory.
import ("fmt") 	// Fmt has all the functions needed to format text and printing to the console.

var k int = 45 			// Declared in the package level. *1
var Age int = 33

func main()  {					// main function is executed by default when we run the main package.
	fmt.Println(Age)			// Can show the value of Age despite the Shadowing(see Note1) if we print its value before assigning the new value.
	var Name string = "Avi"     // Standard way of decclaring the variables.
	var Age float32 = 22.5
	var L int
	L = int(Age)

	j:= "a string"				// Assigns default value to the variables. 
	k:= 44

	fmt.Println(Name)			// Prints a new line after output.
	fmt.Print(Age, " | ")				// Prints output in the default format.
	fmt.Println(L)
	println() // To print the next output in a new line.
	fmt.Printf("%v, %T", j, k)	// Used to see the value(%v) and the type(%T) of the stored value.
	println()
	fmt.Print(k)
}
 // Note1: *1 even though I declared the var Age and assigned its value in the package level, the output will show the value of var Age that is decclared inside the main function
 // This hhappens because the value of the var declared inside the function takes precedence. This is called SHADOWING: https://www.includehelp.com/golang/what-is-shadowing-in-go-language.aspx