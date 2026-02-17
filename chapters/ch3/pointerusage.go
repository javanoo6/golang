package main

import (
	"fmt"
)

func main() {
	amount := 6

	someMethod(&amount)
	fmt.Println(amount)
}

func someMethod(number *int) {
	*number *=  2
}
