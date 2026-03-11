package main

import (
	"fmt"
	"log"

	"chapters/ch5/datafile"
)

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	numbers, err := datafile.GetFloats("datafile/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount = float64(len(numbers))
	fmt.Printf("Average from file: %0.2f\n", sum/sampleCount)
}
