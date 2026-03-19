package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriberOg) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

// возвращает указатель на структуру вместо самой структурыы
func defaultSubscriber(name string) *subscriberOg {
	var s subscriberOg
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

// тк *subscriber это указатель, то будет изменяться ориг
func applyDiscount(s *subscriberOg) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	val := 5
	poiner := &val // poiner хранит адрес val

	fmt.Println(poiner)  // выведет адрес
	fmt.Println(*poiner) // выведет значение по адресу

	*poiner = 10     // изменяем val через указатель
	fmt.Println(val) // 10
}
