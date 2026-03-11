package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
	fmt.Println(notes[3]) // here would be empty string

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[3]) // here would be 0

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	var notesLiteral [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notesLiteral[3], notesLiteral[6], notesLiteral[0])
	var primesLiteral [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primesLiteral[0], primesLiteral[2], primesLiteral[4])

	notesLiteralShort := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notesLiteralShort[3], notesLiteralShort[6], notesLiteralShort[0])
	primesLiteralShort := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primesLiteralShort[0], primesLiteralShort[2], primesLiteralShort[4])
}
