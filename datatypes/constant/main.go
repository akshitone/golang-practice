package main

import "fmt"

const (
	first = iota
	second
	third
)

const (
	newFirst = iota
	newSecond
)

const (
	_ = iota
	startWithOne
)

const (
	_ = iota + 4
	startWithFive
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	const pi = 3.14
	var scopedFirst int
	fmt.Println(pi)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	fmt.Println(newFirst)
	fmt.Println(newSecond)

	fmt.Println(scopedFirst == newFirst)

	fmt.Println(startWithOne)

	fmt.Println(startWithFive)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
