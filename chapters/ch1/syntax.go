package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// fmt.Println("Hello, Go!")//

	fmt.Println(math.Floor(2.75))               // 2
	fmt.Println(strings.Title("head first go")) // Head First Go

	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello go"))

	var quantity int
	var lenght, width float64
	var customerName string

	quantity = 2
	lenght, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(lenght*width, "square meters")
}
