package main

import (
	"fmt"
	"math"
)

func main() {

	// Simple values

	//int
	fmt.Println(1 + 1)

	//string
	fmt.Println("Hello Go Programmer")

	//bool
	fmt.Println(true)
	fmt.Println(false)

	//float
	fmt.Println(5.5)
	fmt.Println(7.5 / 3.2)             //2.34375
	fmt.Println(math.Floor(7.5 / 3.2)) // 2
	fmt.Println(math.Ceil(7.5 / 3.2))  // 3

}
