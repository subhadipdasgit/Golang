package main

import (
	"errors"
	"fmt"
)

func main() {
	var n1, n2 float64
	fmt.Print("Enter first number = ")
	fmt.Scan(&n1)
	fmt.Print("Enter second number = ")
	fmt.Scan(&n2)
	result, error := division(n1, n2)
	if error != nil {
		fmt.Print("Result = ", result, error)
	} else {
		fmt.Print("Result =", result)
	}

}
func division(x, y float64) (float64, error) {
	if x <= 0 || y <= 0 {
		return 0, errors.New("Math error : Zero has been found")
	} else {
		return (x / y), nil
	}
}
