package main

import (
	"fmt"
)

func main() {
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)

	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)
}
