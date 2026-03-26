package main

import (
	"fmt"
	"time"
)

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Println(s)
	}
}

func main() {
	go repeat("x")
	go repeat("y")
	time.Sleep(time.Second)
}
