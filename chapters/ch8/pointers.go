package main

import (
	"fmt"
)

type myStruct struct {
	myField int
}

func main() {
	var value int = 2
	var pointer *int = &value
	fmt.Println("real pointer", pointer)
	fmt.Println("value of pointer", *pointer)

	var mStruct myStruct
	mStruct.myField= 3
	var mPointer *myStruct = &mStruct
	fmt.Println((*mPointer).myField)
	fmt.Println(mPointer.myField)
}
