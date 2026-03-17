package main

import (
	"fmt"
)

func main() {
	notes := make([]string, 7)
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(len(notes))
	fmt.Println(len(primes))
}
