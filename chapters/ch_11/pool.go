package main

import (
	"fmt"
)

type NewTruck string

func (t NewTruck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t NewTruck) Brake() {
	fmt.Println("Stopping")
}

func (t NewTruck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t NewTruck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type NewVehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryNewVehicle(vehicle NewVehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(NewTruck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	TryNewVehicle(NewTruck("Fnord F180"))
}
