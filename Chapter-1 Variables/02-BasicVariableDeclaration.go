//Notes:

/*
Basic Variables
bool: a boolean value, either true or false
string: a sequence of characters
int: a signed integer
float64: a floating-point number
byte: exactly what it sounds like: 8 bits of data
*/

//--------------------------------------------

/*
Question: Assignment
Initialize the variables from the print statement to int, float64, bool and string with their zero values, respectively.
*/

package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var userName string

	fmt.Printf("%d %.2f %v %q\n", smsSendingLimit, costPerSms, hasPermission, userName)
}
