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

	newSubscriber := magazine.Subscriber{Name: "Aman Singh"}
	newSubscriber.HomeAddress.Street = "123 Oak St"
	newSubscriber.HomeAddress.City = "Omaha"
	newSubscriber.HomeAddress.State = "NE"
	newSubscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber Name:", newSubscriber.Name)
	fmt.Println("Street:", newSubscriber.HomeAddress.Street)
	fmt.Println("City:", newSubscriber.HomeAddress.City)
	fmt.Println("State:", newSubscriber.HomeAddress.State)
	fmt.Println("Postal Code:", newSubscriber.HomeAddress.PostalCode)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)
}
