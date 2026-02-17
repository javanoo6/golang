package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

	fmt.Println("-----------------")

	var myInt2 int
	var myIntPointer2 *int
	myIntPointer2 = &myInt2
	fmt.Println(myIntPointer2)
	var myFloat2 float64
	var myFloatPointer2 *float64
	myFloatPointer2 = &myFloat2
	fmt.Println(myFloatPointer2)
}
