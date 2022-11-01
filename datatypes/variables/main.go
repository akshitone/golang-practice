package main

import (
	"fmt"
	"strconv"
)

var owner = "Akshit Mithaiwala" // package scoped variable - lower case first letter
var OWNER = "akshitone"         // global scoped variable - upper case first letter

var (
	johnName          = "John Wick"
	johnAge           = 23
	johnContactNumber = "7874987797"
)

var (
	jennyName          = "Jenny Depend"
	jennyAge           = 21
	jennyContactNumber = "7874977797"
)

func main() {
	var firstVariable int = 34 // block scoped variable
	secondVariable := 23
	fmt.Println(firstVariable + secondVariable)

	var thirdVariable float64
	thirdVariable = float64(firstVariable + secondVariable) // type conversion
	fmt.Println(thirdVariable)

	var converToString string
	converToString = strconv.Itoa(int(thirdVariable)) // int to string conversion
	fmt.Println("The converted string value is " + converToString)

	fmt.Println(owner)
	fmt.Println(OWNER)
}
