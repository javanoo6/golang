package main

import "fmt"

type (
	Liters      float64
	Gallons     float64
	Milliliters float64
)
type Title string

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

// func (g Gallons) ToMilliliters() Milliliters {
// 	return Milliliters(g * 3785.41)
// }

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l* 1000)

}

func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	carFuel := Gallons(Liters(40.0) * 0.264)
	busFuel := Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	fmt.Println(Title("Alien") + "s")
	// fmt.Println(Title("Jaws 2") - " 2")

	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Gallons(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)

	// fmt.Println(Liters(1.2) + Gallons(3.4))
	//	fmt.Println(Gallons(1.2) == Liters(1.2))
	otherCarFule := Gallons(1.2)
	otherBusFuel := Liters(4.5)
	otherCarFule += ToGallons(Liters(40.0))
	otherBusFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", otherCarFule)
	fmt.Printf("Bus fuel: %0.1f liters\n", otherBusFuel)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())

	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", ml, ml.ToLiters())
}
