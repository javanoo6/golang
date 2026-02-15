package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	randomNumber := rand.Intn(100)
	guessCount := 10

	fmt.Println("you need to guess the number between 1 and 100")
	for guessCount >= 0 {

		fmt.Println("you have ", guessCount, " attempts left")
		fmt.Println("guess the number: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guessedNumber, _ := strconv.Atoi(input)

		if guessedNumber == randomNumber {
			fmt.Println("the guess is correct")
			break
		}

		if guessedNumber > randomNumber {
			fmt.Println("you guess is too high")
		} else if guessedNumber < randomNumber {
			fmt.Println("you guess is too low")
		}

		guessCount--
	}
	if guessCount > 0 {
		fmt.Println("you won!")
	} else {
		fmt.Println("you lost, the guessed number was: ", randomNumber)
	}
}
