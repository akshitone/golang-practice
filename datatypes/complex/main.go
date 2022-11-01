package main

import "fmt"

func main() {
	var firstComplexNumber complex64 = 1 + 2i
	fmt.Printf("%T - %v", firstComplexNumber, firstComplexNumber)

	var secondComplexNumber complex64 = 2 + 5.2i

	fmt.Println(firstComplexNumber + secondComplexNumber)
	fmt.Println(firstComplexNumber - secondComplexNumber)
	fmt.Println(firstComplexNumber * secondComplexNumber)
	fmt.Println(firstComplexNumber / secondComplexNumber)
}
