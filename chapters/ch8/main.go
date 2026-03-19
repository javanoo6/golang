package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var subscriber magazine.Subscriber
	subscriber.Name = "Aman Singh"
	subscriber.Rate = 4.99
	subscriber.Active = true
	fmt.Println(subscriber.Name, subscriber.Rate, subscriber.Active)

	subscriberLiteral := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscriberLiteral.Name)
	fmt.Println("Rate:", subscriberLiteral.Rate)
	fmt.Println("Active:", subscriberLiteral.Active)

	var subscriberSemiInit magazine.Subscriber
	subscriberSemiInit.Name = "Aman Singh"
	subscriberSemiInit.Rate = 4.99
	subscriberSemiInit.Active = true
	fmt.Println(subscriberSemiInit.Name, subscriberSemiInit.Rate, subscriberSemiInit.Active)
}
