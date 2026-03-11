package main

import (
	"fmt"
)

func main() {
	var notes [3]string = [3]string{"do", "re", "mi"}
	var primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(notes)
	fmt.Println(primes)

	fmt.Printf("%#v\n", notes) // when using %#v" arrays printed as literals of arrays of Go
	fmt.Printf("%#v\n", primes)

	notesLiteral := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notesLiteral[index])
	index = 3
	fmt.Println(index, notesLiteral[index])

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	// this causes index array out of bounds
	// for i := 0; i <= 7; i++ {
	// 	fmt.Println(i, notesLiteral[i])
	// }

	fmt.Println(len(notes))
	fmt.Println(len(primes))

	// this causes index array out of bounds
	// for i := 0; i <= len(notes); i++ {
	//
	//		fmt.Println(i, notes[i])
	//	}

	// the safe way!!!!
	for index, note := range notesLiteral {
		fmt.Println(index, note)
	}

	for _, note := range notesLiteral { 
		fmt.Println(note) 
	}
}
