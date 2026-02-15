package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Print("Enter grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Print("the input was: ", input)
}
