package main

import "fmt"

func main() {
	var firstVariable bool = true
	fmt.Println(firstVariable)

	secondVariable := false
	fmt.Println(secondVariable)

	var thirdVariable bool
	fmt.Printf("default boolean value: %v \n", thirdVariable)

	isEqual := 1 == 1
	fmt.Println(isEqual)

	isNotEqual := 1 == 2
	fmt.Println(isNotEqual)
}
