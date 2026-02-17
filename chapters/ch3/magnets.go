package main

import "fmt"

func main() {
	var myInt int
	myInt = 42
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}
