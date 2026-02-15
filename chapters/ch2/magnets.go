package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfp, err := os.Stat("some text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfp.Size())
}
