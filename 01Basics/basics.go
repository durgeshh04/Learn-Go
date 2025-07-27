package main

import "fmt"

// variables in Goooo:
/*
Definition:
A variable in Go is a storage location with a name (identifier) and a type, used to hold data that can change during program execution.

Declaration Syntax:
var variableName type = value

var age int = 25           // declares an integer variable
var name string = "GoLang" // declares a string variable
var isActive bool          // declares a boolean variable, default value is false


//Pro tip:-

Use := for quick, local, scoped variables.

Use var for package-level variables or when you need zero values/types explicitly.

Avoid shadowing variables accidentally in large functions.

*/

type Variables struct{} //don't worry about this we've cover this part in our advance file

func (v Variables) integerVariables() int {
	fmt.Println("\nYou've called integer function")
	var a int = 10 // full type declaration
	b := 12        // short hand declaration
	return a + b
}

func (v Variables) stringVariables() (string, rune) {
	fmt.Println("\nYou've called string function")
	var username string = "Durgesh0409"
	var key rune = 'R'
	return username, key
}

// Main function:-
/*
Definition:
In Go, the main function is the entry point of a standalone executable program. It must be defined in the main package. The Go runtime automatically calls main() when you run your program.

Working:

The Go compiler looks for the main package and the main() function.
Execution starts from the first line inside main().
Any code outside main() (like other functions or variable declarations) will not run unless called from main().

*/
func main() {
	fmt.Println("Here We Gooo 🚗...")
	v := Variables{}
	fmt.Println(v.integerVariables())
	name, key := v.stringVariables()
	fmt.Printf("Username: %s, key: %c\n", name, key)
}
