package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	// date := Date{Year: 2019, Month: 5, Day: 27}
	// fmt.Println(date)

	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	// err = date.SetMonth(55)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	event := calendar.Event{}
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
