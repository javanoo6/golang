package main

import (
	"errors"
	"fmt"

)

func devide(dividend float64, divisor float64) (float64, error) {
if divisor == 0.0 {
		return 0, errors.New("cant divide by 0")
	}

	return dividend / divisor, nil
}

func main() {

	qoutient, err :=  devide(5.6, 0.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", qoutient)
	}
}
