package main	// main package is made of the files that are in the same directory.
import ("fmt") 	// Fmt has all the functions needed to format text and printing to the console.

func main()  {					// main function is executed by default when we run the main package.
	var Name string = "Avi"     // Standard way of decclaring the variables.
	var Age int32 = 2205

	j:= "a string"				// Assigns default value to the variables. 
	k:= 44

	fmt.Println(Name)			// Prints a new line after output.
	fmt.Print(Age)				// Prints output in the default format.
	println() // To print the next output in a new line.
	fmt.Printf("%v, %T", j, j)	// Used to see the value(%v) and the type(%T) of the stored value.
	println()
	fmt.Print(k)
}
