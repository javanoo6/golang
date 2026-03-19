package main

import (
	"fmt"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

type part struct {
	description string
	count       int
}
type car struct {
	name     string
	topSpeed float64
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func applyDiscountToCopy(s subscriberOg) {
	s.rate = 4.99
}

func applyDiscountToOrig(s *subscriberOg) {
	s.rate = 4.99
}
type subscriberOg struct {
	name   string
	rate   float64
	active bool
}

func main() {
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count)

	var s subscriberOg
	applyDiscountToCopy(s)
	fmt.Println(s.rate)
	applyDiscountToOrig(&s)
	fmt.Println(s.rate)
}
